package pixicog

import (
  "fmt"
  "testing"
)

func TestMain(t *testing.T) {
  msg, err := run("./test-fixtures/job.go")
  if err != nil {
    t.Fatalf("Unexpected error %v", err)
  }
  exp := "torun -> d68a1f6c58ac03bb8f55ccbaaad6a37445638f14\nHello\n"
  if string(msg) != exp {
    t.Fatalf("Wrong message. Expected [%s]. Got [%s]", exp, msg)
  }
}

func TestBuildMainFunc(t *testing.T) {
  funcs := [][]string{{"run", "sha"}}

  mainFunc, err := buildMainFunc(funcs)
  if err != nil {
    t.Fatalf("Unexpected error %v", err)
  }

  fmt.Println(mainFunc)
}

