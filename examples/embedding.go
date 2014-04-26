package main

import (
	"fmt"
	"github.com/bslatkin/opaque"
)

type MultiplyByTwoType struct {
	opaque.Opaque
}

func (m *MultiplyByTwoType) GetNumber() int {
	return m.Opaque.GetNumber() * 2
}

// Example shows how opaque types can be embedded and methods "overridden".
func main() {
	o := &MultiplyByTwoType{opaque.NewBasic(10)}
	fmt.Printf("My multiply by two value is %d, type is %#v\n", o.GetNumber(), o)

	s := opaque.DoSomethingWithOpaque(o)
	fmt.Println("Doing something with an opaque I get:", s)
}
