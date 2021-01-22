package calendar

import msgraph "github.com/yaegashi/msgraph.go/beta"

// ConvertEvents imports MS Graph events to Event type
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
