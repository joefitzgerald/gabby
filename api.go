package gabby

import "context"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . API
type API interface {
	GetMe(ctx context.Context) (string, error)
	GetRecurringEvents(ctx context.Context) (Events, error)
	GetRecurringEventsWithInstancesForWeeks(ctx context.Context, weeks int) (Events, error)
	GetPhoto(ctx context.Context, id string) (*Photo, error)
	GetName(ctx context.Context, id string) (string, error)
	GetMembers(ctx context.Context, id string) (GroupMembers, error)
}
