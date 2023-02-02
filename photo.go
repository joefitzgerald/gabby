package gabby

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"github.com/llgcode/draw2d/draw2dimg"
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

func RoundRect(path draw2dimg.GraphicContext, x1, y1, x2, y2, arcWidth, arcHeight float64) {
	arcWidth = arcWidth / 2
	arcHeight = arcHeight / 2
	path.MoveTo(x1, y1+arcHeight)
	path.QuadCurveTo(x1, y1, x1+arcWidth, y1)
	path.LineTo(x2-arcWidth, y1)
	path.QuadCurveTo(x2, y1, x2, y1+arcHeight)
	path.LineTo(x2, y2-arcHeight)
	path.QuadCurveTo(x2, y2, x2-arcWidth, y2)
	path.LineTo(x1+arcWidth, y2)
	path.QuadCurveTo(x1, y2, x1, y2-arcHeight)
	path.Close()
}

func (p *Photo) CropRoundRect(width int, arc float64) *Photo {
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
	mask := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	gc := draw2dimg.NewGraphicContext(mask)

	RoundRect(*gc, 3, 3, float64(width-3), float64(width-3), arc, arc)

	gc.SetFillColor(color.Alpha{255})
	gc.Close()
	gc.FillStroke()

	draw.DrawMask(img2, img2.Bounds(), boof, upLeft, mask, upLeft, draw.Over)

	var b bytes.Buffer
	png.Encode(&b, img2)
	return &Photo{
		ID:                 p.ID,
		SuggestedExtension: ".png",
		Data:               b.Bytes(),
	}
}

func (p *Photo) CropCircle(width int) *Photo {
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
