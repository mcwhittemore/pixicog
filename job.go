package pixicog

type Job struct {
  source ImageList
  state ImageList
}

func NewJob(src ImageList) Job {
  j := Job{}
  j.source = src;
  return j
}

func (j Job) GetState() ImageList {
  return j.state
}

func (j Job) Process(fn func(source, state ImageList) ImageList) Job {
  j.state = fn(j.source, j.state)
  return j
}
