package pixicog

import (
  "math"
	"image"
  "image/color"
  "github.com/BurntSushi/graphics-go/graphics"
)

type Pixicog []image.Image

func (p Pixicog) GetFloatPixels(x, y int) FloatPixels {
  fps := make(FloatPixels, len(p))
  model := p.ColorModel()

  for i := 0; i < len(p); i++ {
    c := model.Convert(p[i].At(x, y))
    rt, gt, bt, at := c.RGBA()

    r := float32(uint8(rt))
    g := float32(uint8(gt))
    b := float32(uint8(bt))
    a := float32(uint8(at))

    fps[i] = []float32{r, g, b, a}
  }

  return fps
}

func (p Pixicog) Rotate(deg float64) Pixicog {

  rad := deg * math.Pi / 180

  for i := 0; i < len(p); i++ {
    srcDim := p[i].Bounds()
    dstImage := image.NewRGBA(image.Rect(0, 0, srcDim.Dy(), srcDim.Dx()))
    graphics.Rotate(dstImage, p[i], &graphics.RotateOptions{rad})
    p[i] = dstImage
  }

  return p
}

func (p Pixicog) At(x, y int) color.Color {

  fps := p.GetFloatPixels(x, y)

  var r, g, b, a float32 = 0, 0, 0, 0
  n := float32(len(p))

  for i := 0; i < len(p); i++ {
    r += fps[i][0] / n
    g += fps[i][1] / n
    b += fps[i][2] / n
    a += fps[i][3] / n
  }

  return color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
}

func (p Pixicog) Bounds() image.Rectangle {
  return p[0].Bounds()
}

func (p Pixicog) GetDiminished(x, y, cpc int) []color.Color {

  colors := make([]color.Color, len(p))

  fps := p.GetFloatPixels(x, y)
  cpc8 := uint8(cpc)
  cpcf := float64(cpc8)

  for i := 0; i < len(p); i++ {
    r := uint8(math.Floor(float64(fps[i][0]) / cpcf) * cpcf)
    g := uint8(math.Floor(float64(fps[i][1]) / cpcf) * cpcf)
    b := uint8(math.Floor(float64(fps[i][2]) / cpcf) * cpcf)
    a := uint8(math.Floor(float64(fps[i][3]) / cpcf) * cpcf)
    colors[i] = color.RGBA{r,g,b,a}
  }

  return colors
}

func (p Pixicog) Height() int {
  b := p.Bounds()
  return b.Max.Y - b.Min.Y
}

func (p Pixicog) Width() int {
  b := p.Bounds()
  return b.Max.X - b.Min.X
}

func (p Pixicog) ColorModel() color.Model {
  return p[0].ColorModel()
}
