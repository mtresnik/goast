package goast

import (
	"math/cmplx"
)

type Variable struct {
	Name string
}

func (v Variable) GetValues() []Operation {
	return make([]Operation, 0)
}

func (v Variable) IsConstant() bool {
	return false
}

func (v Variable) String() string {
	return v.Name
}

func (v Variable) Number() complex128 {
	return cmplx.NaN()
}

func (v Variable) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(v, one) {
		return other
	}
	return v
}
