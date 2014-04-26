package main

import "fmt"

type Awesome interface {
	IsAwesome() bool
}

type MyType struct {
	howAwesome int
}

func (m *MyType) IsAwesome() bool {
	return m.howAwesome > 0
}

func PrintItOut(a Awesome) {
	if a.IsAwesome() {
		fmt.Println("This is awesome")
	} else {
		fmt.Println("This is not awesome")
	}
}

func main() {
	PrintItOut(&MyType{0})
	PrintItOut(&MyType{10})
}
