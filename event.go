package gabby

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strings"

	"time"

	"github.com/liggitt/tabwriter"

	"github.com/jinzhu/now"
)

type Events []Event

func (e Events) SortByImpact(start time.Time, weeks int) {
	sort.SliceStable(e, func(i, j int) bool {
		_, totalImpactI := e[i].Impact(start, weeks)
		_, totalImpactJ := e[j].Impact(start, weeks)
		return totalImpactI > totalImpactJ
	})
}

type Event struct {
	ID             string
	Name           string
	Organizer      string
	SeriesMasterID string
	RecurrenceType string
	StartDate      time.Time
	EndDate        time.Time
	Start          time.Time
	End            time.Time
	Interval       int
	Duration       int
	Categories     []string
	Instances      []Event
	Attendees      []Attendee
	Properties     []string
}

func (e *Event) SetProperty(p string) {
	p = strings.ReplaceAll(strings.ToLower(p), " ", "_")
	for _, property := range e.Properties {
		if strings.EqualFold(p, property) {
			return
		}
	}
	e.Properties = append(e.Properties, p)
}

func (e *Event) Impact(start time.Time, weeks int) (result []int, total int) {
	for j := 0; j < weeks; j++ {
		minutes := 0
		weekStart := start.AddDate(0, 0, 7*j)
		weekEnd := now.New(weekStart).EndOfWeek()
		for k := range e.Instances {
			if e.Instances[k].Start.After(weekStart) && e.Instances[k].End.Before(weekEnd) {
				minutes = minutes + e.Instances[k].Duration
			}
		}
		total = total + minutes
		result = append(result, minutes)
	}
	return result, total
}

type Attendee struct {
	Name  string
	Email string
}

func PopulateProperties(events []Event, me string) {
	for i := range events {
		otherAttendees := 0
		for j := range events[i].Attendees {
			if strings.EqualFold(events[i].Attendees[j].Email, me) {
				continue
			}
			otherAttendees = otherAttendees + 1
		}
		if otherAttendees == 0 {
			events[i].SetProperty("self")
		}
		if otherAttendees == 1 {
			events[i].SetProperty("one_on_one")
		}
		for _, category := range events[i].Categories {
			events[i].SetProperty(category)
		}
	}
}

func FilterEvents(events []Event, must []string, mustNot []string) []Event {
	return filter(events, func(e Event) bool {
		if len(must) == 0 && len(mustNot) == 0 {
			return true
		}
		mustSatisfied := true
		for _, m := range must {
			found := false
			for _, property := range e.Properties {
				if strings.EqualFold(m, property) {
					found = true
				}
			}
			if !found {
				mustSatisfied = false
				break
			}
		}
		if !mustSatisfied {
			return false
		}

		for _, m := range mustNot {
			for _, property := range e.Properties {
				if strings.EqualFold(m, property) {
					return false
				}
			}
		}

		return true
	})
}

func filter(events []Event, f func(Event) bool) []Event {
	result := make([]Event, 0)
	for i := range events {
		if f(events[i]) {
			result = append(result, events[i])
		}
	}
	return result
}

func WriteEvents(w io.Writer, events Events, weeks int, detail bool, pretty bool) error {
	start := now.BeginningOfWeek()
	events.SortByImpact(start, weeks)
	tab := '\t'
	var tw *tabwriter.Writer
	var cw *csv.Writer
	if pretty {
		tw = tabwriter.NewWriter(w, 0, 8, 1, byte(tab), tabwriter.TabIndent)
		cw = csv.NewWriter(tw)
	} else {
		cw = csv.NewWriter(w)
	}

	cw.Comma = tab
	header := []string{
		"Name",
		"Organizer",
		"Start Date",
		"End Date",
		"Start",
		"End",
		"Interval",
		"Duration",
		"Categories",
	}
	if detail {
		for i := 0; i < weeks; i++ {
			header = append(header, start.AddDate(0, 0, i*7).Format("2006-01-02"))
		}
	}

	header = append(header, "Impact")
	err := cw.Write(header)
	if err != nil {
		return err
	}
	cw.Flush()
	for i := range events {
		endDate := events[i].EndDate.Local().Format("2006-01-02")
		if events[i].EndDate.IsZero() {
			endDate = "∞"
		}
		name := events[i].Name
		if pretty {
			name = fmt.Sprintf("%.40s", events[i].Name)
		}
		row := []string{
			name,
			events[i].Organizer,
			events[i].StartDate.Local().Format("2006-01-02"),
			endDate,
			events[i].Start.Local().Format(time.Kitchen),
			events[i].End.Local().Format(time.Kitchen),
			fmt.Sprintf("%v", events[i].Interval),
			fmt.Sprintf("%v", events[i].Duration),
			strings.Join(events[i].Categories, ","),
		}
		impactWeeks, impactTotal := events[i].Impact(start, weeks)
		if detail {
			for _, impactWeek := range impactWeeks {
				row = append(row, fmt.Sprintf("%v", impactWeek))
			}
		}

		row = append(row, fmt.Sprintf("%v", impactTotal))
		err = cw.Write(row)
		if err != nil {
			return err
		}
		cw.Flush()
	}
	if pretty {
		return tw.Flush()
	}
	return nil
}
