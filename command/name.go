package command

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/joefitzgerald/gabby"
)

type Name struct {
	IDs []string `arg:"" name:"ids" help:"IDs to retrieve photos for"`
}

func (p *Name) Run(ctx Context, api gabby.API, w io.Writer) error {

	for _, id := range p.IDs {
		name, err := api.GetName(context.Background(), id)
		if err != nil {
			log.Printf("Unable to process name for ID '%s'", id)
			return err
		}

		fmt.Printf("%s, %s\n", id, name)

	}
	return nil
}
