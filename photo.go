package gabby

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/nfnt/resize"
)

type Photo struct {
	Data               []byte
	ID                 string
	SuggestedExtension string
}

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

func (p *Photo) Cropped(width int) *Photo {
	img, _, _ := image.Decode(bytes.NewReader(p.Data))

	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, _ := analyzer.FindBestCrop(img, 100, 100)

	// The crop will have the requested aspect ratio, but you need to copy/scale it yourself

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := img.(SubImager).SubImage(topCrop)
	boof := resize.Resize(uint(width), 0, croppedimg, resize.Lanczos3)

	height := width

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img2 := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	draw.DrawMask(img2, img2.Bounds(), boof, upLeft, &circle{image.Point{width / 2, width / 2}, width / 2}, upLeft, draw.Over)

	var b bytes.Buffer
	png.Encode(&b, img2)
	return &Photo{
		ID:                 p.ID,
		SuggestedExtension: ".png",
		Data:               b.Bytes(),
	}
}

func (p *Photo) SuggestedFilename() string {
	return p.ID + p.SuggestedExtension
}
