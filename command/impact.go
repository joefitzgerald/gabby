package command

import (
	"context"
	"io"

	"github.com/joefitzgerald/gabby"
)

type Impact struct {
	Weeks   int      `name:"weeks" short:"w" default:"8" help:"The number of weeks to include in the impact analysis."`
	Include []string `name:"include" short:"i" optional:"" help:"Properties that must be present for an event."`
	Exclude []string `name:"exclude" short:"e" optional:"" help:"Properties that must not be present for an event."`
	Detail  bool     `name:"detail" optional:"" default:"false" help:"Include detail showing the minutes of impact for each week."`
	Pretty  bool     `name:"pretty" optional:"" default:"true" help:"Print output so that it displays nicely in the terminal. Disable to pipe to pbcopy and import to Google sheets."`
}

func (i *Impact) Run(ctx Context, api gabby.API, w io.Writer) error {
	events, err := api.GetRecurringEventsWithInstancesForWeeks(context.Background(), 8)
	if err != nil {
		return err
	}
	gabby.PopulateProperties(events, ctx.Me)
	events = gabby.FilterEvents(events, i.Include, i.Exclude)
	return gabby.WriteEvents(w, events, i.Weeks, i.Detail, i.Pretty)
}
