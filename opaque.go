package opaque

import "fmt"

type Opaque interface {
	GetNumber() int
	nobodyOutsideCanImplementThis()
}

type internalBasicType int

func (m *internalBasicType) GetNumber() int {
	return int(*m)
}

func (m *internalBasicType) nobodyOutsideCanImplementThis() {}

func NewBasic(number int) Opaque {
	o := internalBasicType(number)
	return &o
}

type internalStructType struct {
	number int
}

func (m *internalStructType) GetNumber() int {
	return m.number
}

func (m *internalStructType) nobodyOutsideCanImplementThis() {}

func NewStruct(number int) Opaque {
	return &internalStructType{number}
}

func DoSomethingWithOpaque(o Opaque) string {
	return fmt.Sprintf("Hello opaque #%d", o.GetNumber())
}
