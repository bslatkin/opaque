package main

import (
	"fmt"
	"github.com/bslatkin/opaque"
)

func main() {
	o := opaque.NewBasic(15)
	fmt.Printf("My basic value is %d\n", o.GetNumber())

	o = opaque.NewStruct(15)
	fmt.Printf("My struct value is %d\n", o.GetNumber())
}
