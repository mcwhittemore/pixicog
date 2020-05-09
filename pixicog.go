package pixicog

import (
	"image"
  "image/color"
)

type Pixicog []*image.RGBA

func (p Pixicog) At(x, y int) color.Color {
  model := p.ColorModel()

  var r, g, b, a uint8 = 0, 0, 0, 0
  n := uint8(len(p))

  count := 0
  for i := 0; i < len(p); i++ {
    count++
    c := model.Convert(p[i].At(x, y))
    rt, gt, bt, at := c.RGBA()
    r += uint8(rt) / n
    g += uint8(gt) / n
    b += uint8(bt) / n
    a += uint8(at) / n
  }

  return color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
}

func (p Pixicog) Bounds() image.Rectangle {
  return p[0].Bounds()
}

func (p Pixicog) ColorModel() color.Model {
  return p[0].ColorModel()
}
