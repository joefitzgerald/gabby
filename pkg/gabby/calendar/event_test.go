package calendar_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/joefitzgerald/gabby/pkg/gabby/calendar"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testCalendar(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	it("", func() {
		b := &bytes.Buffer{}
		now := time.Now()
		events := []Event{
			{
				"Test",
				"noEnd",
				now,
				now,
			},
		}
		err := calendar.WriteEvents(b, events)
		Expect(err).NotTo(HaveOccurred())
		expected := []string{
			"Name",
			"Recurrence Type",
			"Start Date",
			"End Date",
		}
		Expect(b.String()).To(Equal(strings.Join(expected, "\t") + "\n"))
	})
}
