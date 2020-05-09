package pixicog

import (
  "os"
  "log"
  "testing"
  "image"
  "image/png"
  "image/color"
)

func TestPixicogIsAnImageInterface(t *testing.T) {
  cog := Pixicog{}
  white := color.RGBA{255, 255, 255, 255};
  black := color.RGBA{0, 0, 0, 255};

  cog = append(cog, FlatImage(10, 10, white)) // white
  cog = append(cog, FlatImage(10, 10, black)) // black

  r, g, b, a := cog.At(0,0).RGBA()
  if uint8(r) != 127 {
		t.Fatalf("Red is incorrect. Expected 127 but got %d", uint8(r))
  }
  if uint8(g) != 127 {
		t.Fatalf("Green is incorrect. Expected 127 but got %d", uint8(g))
  }
  if uint8(b) != 127 {
		t.Fatalf("Blue is incorrect. Expected 127 but got %d", uint8(b))
  }
  if uint8(a) != 255 {
		t.Fatalf("Alpha is incorrect. Expected 255 but got %d", uint8(a))
  }

  Save(cog, t)
}

func FlatImage(width, height int, c color.RGBA) *image.RGBA {
  img := image.NewRGBA(image.Rect(0,0,width,height))

  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      img.Set(x, y, c)
    }
  }

  return img
}

func Save(img image.Image, t *testing.T) {
  filename := "./test-artifacts/" + t.Name() + ".png"
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

