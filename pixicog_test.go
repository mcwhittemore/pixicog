package pixicog

import (
  "os"
  "log"
  "testing"
  "image"
  "image/png"
  "image/color"
)

func TestFloatPixel(t *testing.T) {
  cog := Pixicog{}
  white := color.RGBA{255, 255, 255, 255};
  black := color.RGBA{0, 0, 0, 255};
  gray := color.RGBA{128, 128, 128, 255};

  cog = append(cog, FlatImage(1, 1, white)) // white
  cog = append(cog, FlatImage(1, 1, black)) // black
  cog = append(cog, FlatImage(1, 1, gray)) // gray

  rw, _, _, _ := cog.FloatPixel(0, 0, 0)
  rb, _, _, _ := cog.FloatPixel(1, 0, 0)
  rg, _, _, _ := cog.FloatPixel(2, 0, 0)

  if rw != 255 {
		t.Fatalf("White is incorrect. Expected 255 but got %f", rw)
  }

  if rb != 0 {
		t.Fatalf("Black is incorrect. Expected 0 but got %f", rb)
  }

  if rg != 128 {
		t.Fatalf("Gray is incorrect. Expected 128 but got %f", rg)
  }

}

func TestGetDiminishedReturnsExpectedColor(t * testing.T) {
  cog := Pixicog{}
  white := color.RGBA{255, 255, 255, 255};
  black := color.RGBA{0, 0, 0, 255};
  gray := color.RGBA{128, 128, 128, 255};

  cog = append(cog, FlatImage(1, 1, white)) // white
  cog = append(cog, FlatImage(1, 1, black)) // black
  cog = append(cog, FlatImage(1, 1, gray)) // gray

  colors := cog.GetDiminished(0,0,16)
  wr, _, _, _ := colors[0].RGBA()
  br, _, _, _ := colors[1].RGBA()
  gr, _, _, _ := colors[2].RGBA()

  if uint8(wr) != 240 {
		t.Fatalf("White is incorrect. Expected 240 but got %d", uint8(wr))
  }

  if uint8(br) != 0 {
		t.Fatalf("Black is incorrect. Expected 0 but got %d", uint8(br))
  }

  if uint8(gr) != 128 {
		t.Fatalf("Gray is incorrect. Expected 128 but got %d", uint8(gr))
  }
}

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

