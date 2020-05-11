package pixicog

import (
  "image/color"
)

type FloatPixelList []FloatPixel

func (fps FloatPixelList) GetColors(fn func(FloatPixel) FloatPixel) []color.Color {
  colors := make([]color.Color, len(fps))

  for i := 0; i < len(fps); i++ {
    colors[i] = fps[i].GetColor(fn)
  }

  return colors
}

func (fps FloatPixelList) GetColor(fn func(FloatPixel, FloatPixel) FloatPixel) color.Color {
  fp := make(FloatPixel, len(fps[0]))
  n := len(fps)

  for i := 0; i < n; i++ {
    fp = fn(fp, fps[i])
  }

  return fp.GetColor(nil)
}
