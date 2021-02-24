package gabby_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/joefitzgerald/gabby"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testCalendar(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	it("writes tab separated output with a header", func() {
		b := &bytes.Buffer{}
		now := time.Now()
		events := []gabby.Event{
			{
				Name:           "Test",
				Organizer:      "t",
				RecurrenceType: "noEnd",
				StartDate:      now,
				EndDate:        now,
				Start:          now,
				End:            now,
				Interval:       60,
			},
		}
		err := gabby.WriteEvents(b, events, 1, false, false)
		Expect(err).NotTo(HaveOccurred())
		expectedHeader := []string{
			"Name",
			"Organizer",
		}
		Expect(b.String()).To(gomega.HavePrefix(strings.Join(expectedHeader, "\t")))
	})

	when("filtering events", func() {
		var events []gabby.Event
		it.Before(func() {
			events = []gabby.Event{
				{
					Name:       "1",
					Properties: []string{"self", "one_on_one", "recurring", "private"},
				},
				{
					Name:       "2",
					Properties: []string{"self", "recurring", "public"},
				},
				{
					Name:       "3",
					Properties: []string{"recurring"},
				},
			}
		})

		it("includes all events when must and must not are empty", func() {
			actual := gabby.FilterEvents(events, nil, nil)
			Expect(actual).To(HaveLen(len(events)))
		})

		it("includes events that match a must argument with one element", func() {
			actual := gabby.FilterEvents(events, []string{"self"}, nil)
			Expect(actual).To(HaveLen(2))
			actual = gabby.FilterEvents(events, []string{"private"}, nil)
			Expect(actual).To(HaveLen(1))
		})

		it("excludes events that match a must not argument with one element", func() {
			actual := gabby.FilterEvents(events, nil, []string{"recurring"})
			Expect(actual).To(HaveLen(0))
			actual = gabby.FilterEvents(events, nil, []string{"one_on_one"})
			Expect(actual).To(HaveLen(2))
		})

		it("includes events that match an include but not the exclude list", func() {
			actual := gabby.FilterEvents(events, []string{"recurring"}, []string{"self"})
			Expect(actual).To(HaveLen(1))
		})
	})
}
