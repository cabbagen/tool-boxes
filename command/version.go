package command

import "fmt"

type version struct {
  name          string
  value         string
  description   string
}

func (v version) excute() {
  fmt.Printf("The version of %s: %s\n", v.name, v.value);
}

func NewVersion() version {
  return version {
    name: "tool-boxes",
    value: "1.0.0",
    description: "This is tools collection by golang",
  };
}
