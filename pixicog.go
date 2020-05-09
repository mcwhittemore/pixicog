package pixicog

import (
	"image"
  "image/color"
)

type Pixicog []*image.RGBA

func (p Pixicog) At(x, y int) color.Color {
  model := p.ColorModel()

  var r, g, b, a uint8 = 0, 0, 0, 0
  n := uint32(len(p)) + 1

  for i := uint32(0); i < n - 1; i++ {
    c := model.Convert(p[i].At(x, y))
    rt, gt, bt, at := c.RGBA()
    r += uint8(rt)
    g += uint8(gt)
    b += uint8(bt)
    a += uint8(at)

  }

  r = r / n
  g = g / n
  b = b / n
  a = a / n

  return color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
}

func (p Pixicog) Bounds() image.Rectangle {
  return p[0].Bounds()
}

func (p Pixicog) ColorModel() color.Model {
  return p[0].ColorModel()
}
