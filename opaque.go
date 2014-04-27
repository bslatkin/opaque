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

func (m *internalBasicType) nobodyOutsideCanImplementThis() {
	fmt.Println("Called nobodyOutsideCanImplementThis")
}

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

func VerifyAndDoSomethingWithOpaque(o Opaque) (err error, result string) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = recovered.(error)
		}
	}()
	o.nobodyOutsideCanImplementThis()
	return nil, DoSomethingWithOpaque(o)
}
