package pixicog

import (
  "image/color"
  "testing"
)

func TestJobProcessesWorks(t *testing.T) {
  cog := ImageList{}
  blue := color.RGBA{0, 0, 255, 255};
  green := color.RGBA{0, 255, 0, 255};

  job := NewJob(cog)
  job = job.Process(func(source, state ImageList) ImageList {
    return source // moves source to state
  }).Process(func(source, state ImageList) ImageList {
    return append(state, FlatImage(10, 10, blue))
  }).Process(func(source, state ImageList) ImageList {
    return append(state, FlatImage(10, 10, green))
  })

  Save(job.GetState(), t)
}
