package main

import (
	"fmt"
	"github.com/bslatkin/opaque"
)

type MultiplyByX struct {
	opaque.Opaque
	x int
}

func (m *MultiplyByX) GetNumber() int {
	return m.Opaque.GetNumber() * m.x
}

// Example trying to show how to cause a runtime panic with calls to a hidden
// method on a public interface.
func main() {
	o := &MultiplyByX{Opaque: opaque.NewBasic(10), x: 3}
	fmt.Printf("My multiply by X value is %d, type is %#v\n", o.GetNumber(), o)

	s := opaque.DoSomethingWithOpaque(o)
	fmt.Println("Doing something with an opaque I get:", s)

	err, s := opaque.VerifyAndDoSomethingWithOpaque(o)
	fmt.Printf("Verifying %#v I get: %#v, %#v\n", o, err, s)

	o = &MultiplyByX{Opaque: nil, x: 3}
	err, s = opaque.VerifyAndDoSomethingWithOpaque(o)
	fmt.Printf("Verifying %#v I get: %#v, %#v\n", o, err, s)
}
