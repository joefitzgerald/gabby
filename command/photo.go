package command

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joefitzgerald/gabby"
)

type Photo struct {
	FromFile        *os.File `help:"read ids from a file, specify stdin with -"`
	IDs             []string `arg:"" name:"ids" help:"IDs to retrieve photos for" optional:""`
	OutputDirectory string   `type:"path" default:"./"`
	CropCircle      bool     `default:"false"`
	CropWidth       int      `default:"200"`
}

func (p *Photo) Run(ctx Context, api gabby.API, w io.Writer) error {
	var ids []string

	ids = append(ids, p.IDs...)

	if p.FromFile != nil {
		scanner := bufio.NewScanner(p.FromFile)
		for scanner.Scan() {
			sep := strings.Split(scanner.Text(), ",")
			for i := range sep {
				id := strings.TrimSpace(sep[i])
				if len(id) > 0 {
					ids = append(ids, id)
				}
			}
		}
	}

	for _, id := range ids {
		photo, err := api.GetPhoto(context.Background(), id)
		if err != nil {
			log.Printf("Unable to process photo for ID '%s'", id)
		} else {

			if p.CropCircle {
				photo = photo.CropCircle(p.CropWidth)
			}

			outPath := filepath.Join(p.OutputDirectory, photo.SuggestedFilename())

			if err := os.WriteFile(outPath, photo.Data, 0666); err != nil {
				log.Fatalf("Unable to write file '%s'", outPath)

			}
		}
	}
	return nil
}
