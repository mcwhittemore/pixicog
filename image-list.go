package pixicog

import (
  "math"
  "os"
	"image"
  "image/png"
  "log"
  "image/color"
  "github.com/BurntSushi/graphics-go/graphics"
)

type ImageList []image.Image

func (img ImageList) SavePNG(filename string) {
  f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (p ImageList) GetFloatPixelList(x, y int) FloatPixelList {
  fps := make(FloatPixelList, len(p))
  model := p.ColorModel()

  for i := 0; i < len(p); i++ {
    c := model.Convert(p[i].At(x, y))
    fps[i] = NewFloatPixel(c)
  }

  return fps
}

func (p ImageList) Rotate(deg float64) ImageList {
  out := make(ImageList, len(p))

  rad := deg * math.Pi / 180

  for i := 0; i < len(p); i++ {
    srcDim := p[i].Bounds()
    dstImage := image.NewRGBA(image.Rect(0, 0, srcDim.Dy(), srcDim.Dx()))
    graphics.Rotate(dstImage, p[i], &graphics.RotateOptions{rad})
    out[i] = dstImage
  }

  return out
}

func (p ImageList) At(x, y int) color.Color {

  fps := p.GetFloatPixelList(x, y)
  n := float32(len(p))

  scale := func(v float32, idx int) float32 {
    return v / n
  }

  merge := func(a, b FloatPixel) FloatPixel {
    b = b.Map(scale)
    return a.Add(b)
  }

  return fps.GetColor(merge)
}

func (p ImageList) Bounds() image.Rectangle {
  return p[0].Bounds()
}

func (p ImageList) GetDiminished(x, y, cpc int) []color.Color {
  fps := p.GetFloatPixelList(x, y)
  cpc8 := uint8(cpc)
  cpcf := float64(cpc8)

  each := func(v float32, idx int) float32 {
    return float32(math.Floor(float64(v) / cpcf) * cpcf)
  }

  fn := func(fp FloatPixel) FloatPixel {
    return fp.Map(each)
  }

  return fps.GetColors(fn)
}

func (p ImageList) Height() int {
  b := p.Bounds()
  return b.Max.Y - b.Min.Y
}

func (p ImageList) Width() int {
  b := p.Bounds()
  return b.Max.X - b.Min.X
}

func (p ImageList) ColorModel() color.Model {
  return p[0].ColorModel()
}
