package main

import (
	"fmt"
	"github.com/bslatkin/opaque"
)

// Simple example shows how you can use a package with exported, opaque types.
func main() {
	o := opaque.NewBasic(15)
	fmt.Printf("My basic value is %d, type is %#v\n", o.GetNumber(), o)

	o = opaque.NewStruct(33)
	fmt.Printf("My struct value is %d, type is %#v\n", o.GetNumber(), o)
}
