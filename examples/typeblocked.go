package main

import (
	"fmt"
	"github.com/bslatkin/opaque"
)

type TryingToBeOpaque struct {
	count int
}

func (m *TryingToBeOpaque) GetNumber() int {
	return m.count
}

// Required to implement the Opaque interface, but not allowed because the
// method is not exported from the opaque package. This fails to compile.
func (m *TryingToBeOpaque) nobodyOutsideCanImplementThis() {}

func main() {
	o := &TryingToBeOpaque{10}
	s := opaque.DoSomethingWithOpaque(o)
	fmt.Println("Doing something with an opaque I get:", s)
}
