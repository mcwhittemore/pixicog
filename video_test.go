package pixicog

import "testing"

func TestPixicogFromVideoFileNameToReturnAPixicog(t *testing.T) {
	cog, err := PixicogFromVideoFileName("./test-fixtures/gmf-test-video.mp4")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	actualNumImages := len(cog)
	expectedNumImages := 25
	if actualNumImages != expectedNumImages {
		t.Fatalf("Expected %d images but got %d", expectedNumImages, actualNumImages)
	}
}
