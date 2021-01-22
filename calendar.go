package gabby

import (
	"encoding/csv"
	"io"
	"time"
	"unicode/utf8"

	msgraph "github.com/yaegashi/msgraph.go/beta"
)

type Event struct {
	Name      string
	Type      string
	StartDate time.Time
	EndDate   time.Time
}

func ConvertEvents(events []msgraph.Event) []Event {
	result := make([]Event, len(events))
	for i := range events {
		e := Event{}
		if events[i].Subject != nil {
			e.Name = *events[i].Subject
		}

		if events[i].Recurrence != nil && events[i].Recurrence.Range != nil {
			if events[i].Recurrence.Range.Type != nil {
				e.Type = string(*events[i].Recurrence.Range.Type)
			}
		}

		if events[i].Recurrence.Range != nil {
			if events[i].Recurrence.Range.StartDate != nil {
				t, err := events[i].Recurrence.Range.StartDate.Time()
				if err != nil {
					e.StartDate = t
				}
			}
			if events[i].Recurrence.Range.EndDate != nil {
				t, err := events[i].Recurrence.Range.EndDate.Time()
				if err != nil {
					e.EndDate = t
				}
			}
		}
		result = append(result, e)
	}
	return result
}

func WriteEvents(w io.Writer, events []Event) error {
	cw := csv.NewWriter(w)
	tab, _ := utf8.DecodeRuneInString("\t")
	cw.Comma = tab
	header := []string{
		"Name",
		"Recurrence Type",
		"Start Date",
		"End Date",
	}
	for i := range events {
		cw.Write([]string{
			events[i].Name,
			events[i].Type,
			events[i].StartDate.Format(time.RFC3339),
			events[i].EndDate.Format(time.RFC3339),
		})
	}
	defer cw.Flush()
	return cw.Write(header)
}
