package pixicog

import (
  "math"
	"image"
  "image/color"
)

type Pixicog []*image.RGBA

func (p Pixicog) FloatPixel(i, x, y int) (float32, float32, float32, float32) {
  model := p.ColorModel()
  c := model.Convert(p[i].At(x, y))
  rt, gt, bt, at := c.RGBA()

  // while rt, gt, bt, and at are all uint32 numbers converting them
  // straight to a float results in larger than desired numbers and
  // so first we cast them to uint8s as these are all 0 - 255 numbers
  // and than recast it to uint32 before converting into float32s.
  // My guess is that the color convert function doesn't clean up
  // the converted colors uint32s correctly

  r := math.Float32frombits(uint32(uint8(rt)))
  g := math.Float32frombits(uint32(uint8(gt)))
  b := math.Float32frombits(uint32(uint8(bt)))
  a := math.Float32frombits(uint32(uint8(at)))

  return r, g, b, a
}

func (p Pixicog) At(x, y int) color.Color {

  var r, g, b, a float32 = 0, 0, 0, 0
  n := math.Float32frombits(uint32(len(p)))

  for i := 0; i < len(p); i++ {
    rt, gt, bt, at := p.FloatPixel(i, x, y)

    r += rt / n
    g += gt / n
    b += bt / n
    a += at / n
  }

  return color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
}

func (p Pixicog) Bounds() image.Rectangle {
  return p[0].Bounds()
}

func (p Pixicog) ColorModel() color.Model {
  return p[0].ColorModel()
}
