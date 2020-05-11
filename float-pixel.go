package pixicog

import (
  "image/color"
)

type FloatPixel []float32

func NewFloatPixel(c color.Color) FloatPixel {
  fp := make(FloatPixel, 4)

  rt, gt, bt, at := c.RGBA()

  fp[0] = float32(uint8(rt))
  fp[1] = float32(uint8(gt))
  fp[2] = float32(uint8(bt))
  fp[3] = float32(uint8(at))

  return fp
}

func (fp FloatPixel) GetColor(fn func(FloatPixel) FloatPixel) color.Color {
  if fn != nil {
    fp = fn(fp)
  }

  return color.RGBA{uint8(fp[0]), uint8(fp[1]), uint8(fp[2]), uint8(fp[3])}
}

func (fp FloatPixel) Map(fn func(float32, int) float32) FloatPixel {
  nfp := make(FloatPixel, len(fp))
  n := len(fp)
  for i := 0; i < n; i++ {
    nfp[i] = fn(fp[i], i)
  }
  return nfp
}

func (fp FloatPixel) Add(b FloatPixel) FloatPixel {
  o := make(FloatPixel, len(fp))
  n := len(fp)
  for i := 0; i < n; i++ {
    o[i] = fp[i] + b[i]
  }
  return o
}

type FloatPixels []FloatPixel

func (fps FloatPixels) GetColors(fn func(FloatPixel) FloatPixel) []color.Color {
  colors := make([]color.Color, len(fps))

  for i := 0; i < len(fps); i++ {
    colors[i] = fps[i].GetColor(fn)
  }

  return colors
}

func (fps FloatPixels) GetColor(fn func(FloatPixel, FloatPixel) FloatPixel) color.Color {
  fp := make(FloatPixel, len(fps[0]))
  n := len(fps)

  for i := 0; i < n; i++ {
    fp = fn(fp, fps[i])
  }

  return fp.GetColor(nil)
}


