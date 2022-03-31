package msgraph

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/joefitzgerald/gabby"
)

func (a *API) GetMembers(ctx context.Context, id string) (gabby.GroupMembers, error) {
	log.Printf("%s", a.Client.Groups().ID(id).Members().Request().URL())
	members, err := a.Client.Groups().ID(id).Members().Request().Get(ctx)

	if err != nil {

		log.Println("unable to get group member details")
		return nil, err
	}

	// var groupMembers gabby.GroupMembers

	spew.Dump(members)
	// for i := range  {
	// 	f := group.Members[i]
	// 	spew.Dump(f)
	// }
	return nil, nil
}
