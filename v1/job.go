package pixicog

type Job struct {
  source Pixicog
  state Pixicog
}

func NewJob(src Pixicog) Job {
  j := Job{}
  j.source = src;
  return j
}

func (j Job) GetState() Pixicog {
  return j.state
}

func (j Job) Process(fn func(source, state Pixicog) Pixicog) Job {
  j.state = fn(j.source, j.state)
  return j
}
