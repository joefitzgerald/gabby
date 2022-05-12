package command

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/joefitzgerald/gabby"
)

type Photo struct {
	IDs             []string `arg:"" name:"ids" help:"IDs to retrieve photos for"`
	OutputDirectory string   `type:"path" default:"./"`
	CropCircle      bool     `default:"false"`
	CropWidth       int      `default:"200"`
}

func (p *Photo) Run(ctx Context, api gabby.API, w io.Writer) error {

	for _, id := range p.IDs {
		photo, err := api.GetPhoto(context.Background(), id)
		if err != nil {
			log.Printf("Unable to process photo for ID '%s'", id)
			return err
		}

		if p.CropCircle {
			photo = photo.CropCircle(p.CropWidth)
		}

		outPath := filepath.Join(p.OutputDirectory, photo.SuggestedFilename())

		if err := os.WriteFile(outPath, photo.Data, 0666); err != nil {
			log.Fatalf("Unable to write file '%s'", outPath)

		}
	}
	return nil
}
