package gabby

import (
	"time"

	"github.com/jinzhu/now"
)

func RangeForWeeks(n int) (time.Time, time.Time) {
	return now.BeginningOfWeek(), now.New(time.Now().AddDate(0, 0, 7*(n-1))).EndOfWeek()
}
