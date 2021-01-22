package gabby_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

var suite spec.Suite

func init() {
	suite = spec.New("Gabby", spec.Report(report.Terminal{}))
	suite("calendar", testCalendar)
}

func TestGabby(t *testing.T) {
	suite.Run(t)
}
