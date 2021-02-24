package msgraph

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/joefitzgerald/gabby"
	msgraph "github.com/yaegashi/msgraph.go/beta"
)

const msgraphDateFormat string = "2006-01-02T15:04:05.0000000"

func ConvertEvents(events []msgraph.Event) gabby.Events {
	var result gabby.Events
	for i := range events {
		e := gabby.Event{}
		if events[i].ID == nil {
			continue
		}
		e.ID = *events[i].ID
		if events[i].Subject != nil {
			e.Name = *events[i].Subject
		}

		if events[i].SeriesMasterID != nil {
			e.SeriesMasterID = *events[i].SeriesMasterID
		}

		if events[i].Organizer != nil && events[i].Organizer.EmailAddress != nil && events[i].Organizer.EmailAddress.Address != nil {
			e.Organizer = *events[i].Organizer.EmailAddress.Address
		}

		if events[i].Recurrence != nil && events[i].Recurrence.Range != nil {
			e.SetProperty("recurring")
			if events[i].Recurrence.Range.Type != nil {
				e.RecurrenceType = string(*events[i].Recurrence.Range.Type)
			}
		}

		if events[i].Recurrence != nil && events[i].Recurrence.Pattern != nil {
			if events[i].Recurrence.Pattern.Interval != nil {
				e.Interval = *events[i].Recurrence.Pattern.Interval
			}
		}

		e.Categories = events[i].Categories

		if events[i].Start != nil {
			t, err := time.Parse(msgraphDateFormat, *events[i].Start.DateTime)
			if err == nil {
				e.Start = t
			} else {
				log.Fatal(err)
			}
		}

		if events[i].End != nil {
			t, err := time.Parse(msgraphDateFormat, *events[i].End.DateTime)
			if err == nil {
				e.End = t
			}
		}

		if events[i].Recurrence != nil && events[i].Recurrence.Range != nil {
			if events[i].Recurrence.Range.StartDate != nil {
				t, err := events[i].Recurrence.Range.StartDate.Time()
				if err == nil {
					e.StartDate = t
				}
			}
			if events[i].Recurrence.Range.EndDate != nil {
				t, err := events[i].Recurrence.Range.EndDate.Time()
				if err == nil {
					e.EndDate = t
				}
			}
		}
		if !e.Start.IsZero() && !e.End.IsZero() {
			e.Duration = int(e.End.Sub(e.Start).Minutes())
		}
		for j := range events[i].Attendees {
			if events[i].Attendees[j].EmailAddress == nil {
				continue
			}
			attendee := gabby.Attendee{
				Email: *events[i].Attendees[j].EmailAddress.Address,
				Name:  *events[i].Attendees[j].EmailAddress.Name,
			}
			e.Attendees = append(e.Attendees, attendee)
		}

		if events[i].IsAllDay != nil && *events[i].IsAllDay {
			e.SetProperty("all_day")
		}
		if events[i].IsOrganizer != nil && *events[i].IsOrganizer {
			e.SetProperty("organizer")
		}
		if events[i].IsCancelled != nil && *events[i].IsCancelled {
			e.SetProperty("cancelled")
		}
		if events[i].IsOnlineMeeting != nil && *events[i].IsOnlineMeeting {
			e.SetProperty("online_meeting")
		}

		result = append(result, e)
	}
	return result
}

func readEventCache(filename string) ([]msgraph.Event, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var events []msgraph.Event
	err = json.Unmarshal(b, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func writeEventCache(filename string, events []msgraph.Event) error {
	b, err := json.Marshal(events)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0644)
}

func (a *API) getEvents(ctx context.Context) ([]msgraph.Event, error) {
	events, err := readEventCache("events.json")
	if err != nil {
		events, err = a.Client.Me().Calendar().Events().Request().Get(ctx)
		if err != nil {
			return nil, err
		}
		err = writeEventCache("events.json", events)
		if err != nil {
			return nil, err
		}
	}
	return events, nil
}

func (a *API) getEventInstances(ctx context.Context, id string, start time.Time, end time.Time) ([]msgraph.Event, error) {
	req := a.Client.Me().Events().ID(id).Instances().Request()
	req.Query().Add("startDateTime", start.Format(msgraphDateFormat))
	req.Query().Add("endDateTime", end.Format(msgraphDateFormat))
	return req.Get(ctx)
}

func (a *API) GetRecurringEventsWithInstancesForWeeks(ctx context.Context, weeks int) (gabby.Events, error) {
	events, err := a.GetRecurringEvents(ctx)
	if err != nil {
		return nil, err
	}
	start, end := gabby.RangeForWeeks(weeks)
	for i := range events {
		instances, err := a.getEventInstances(ctx, events[i].ID, start, end)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		events[i].Instances = ConvertEvents(instances)
	}
	return events, nil
}

func (a *API) GetRecurringEvents(ctx context.Context) (gabby.Events, error) {
	events, err := a.getEvents(ctx)
	if err != nil {
		return nil, err
	}

	var recurring []msgraph.Event
	for i := range events {
		if events[i].Recurrence == nil {
			continue
		}
		if events[i].Recurrence.Range == nil {
			continue
		}
		if events[i].Recurrence.Range.EndDate == nil || *events[i].Recurrence.Range.Type == "noEnd" {
			recurring = append(recurring, events[i])
			continue
		}
		t, err := events[i].Recurrence.Range.EndDate.Time()
		if err != nil {
			// TODO: is this a valid event to include?
			continue
		}
		if t.After(time.Now()) {
			recurring = append(recurring, events[i])
		}
	}
	return ConvertEvents(recurring), nil
}
