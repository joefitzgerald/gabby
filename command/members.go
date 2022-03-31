package command

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/joefitzgerald/gabby"
)

type Members struct {
	ID string `arg:"" name:"id" help:"ID of group"`
}

func (p *Members) Run(ctx Context, api gabby.API, w io.Writer) error {

	members, err := api.GetMembers(context.Background(), p.ID)
	if err != nil {
		log.Printf("Unable to process groups for ID '%s'", p.ID)
		return err
	}

	fmt.Printf("%s\n", members)

	return nil
}
