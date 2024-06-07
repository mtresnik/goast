package functions

import (
	"goast/pkg/operations"
	"goast/pkg/utils"
	"math/cmplx"
)

type Abs struct {
	Inner operations.Operation
}

func (a Abs) GetValues() []operations.Operation {
	return []operations.Operation{a.Inner}
}

func (a Abs) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a Abs) ToString() string {
	if a.Inner.IsConstant() {
		c := a.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := "abs("
	retString += a.Inner.ToString()
	retString += ")"
	return retString
}

func (a Abs) ToNumber() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return complex(cmplx.Abs(a.Inner.ToNumber()), 0)
}

func (a Abs) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Abs{
		Inner: a.Inner.Evaluate(one, other),
	}
}
