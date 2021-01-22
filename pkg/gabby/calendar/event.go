package calendar

import (
	"encoding/csv"
	"io"
	"time"
	"unicode/utf8"
)

// Event Mapping
type Event struct {
	Name      string
	Type      string
	StartDate time.Time
	EndDate   time.Time
}

// WriteEvents creates a TSV of all events
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
