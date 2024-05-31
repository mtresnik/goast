package variables

import (
	"goast/pkg/operations"
	"math/cmplx"
)

type Variable struct {
	Name string
}

func (v Variable) GetValues() []operations.Operation {
	return make([]operations.Operation, 0)
}

func (v Variable) IsConstant() bool {
	return false
}

func (v Variable) ToString() string {
	return v.Name
}

func (v Variable) ToNumber() complex128 {
	return cmplx.NaN()
}

func (v Variable) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(v, one) {
		return other
	}
	return v
}
