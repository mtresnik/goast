package functions

import (
	"goast/pkg/operations"
	"goast/pkg/utils"
	"math/cmplx"
)

type ArcSin struct {
	Inner operations.Operation
}

func (a ArcSin) GetValues() []operations.Operation {
	return []operations.Operation{a.Inner}
}

func (a ArcSin) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a ArcSin) ToString() string {
	if a.Inner.IsConstant() {
		c := a.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := "arcsin("
	retString += a.Inner.ToString()
	retString += ")"
	return retString
}

func (a ArcSin) ToNumber() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Asin(a.Inner.ToNumber())
}

func (a ArcSin) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return ArcSin{Inner: a.Inner.Evaluate(one, other)}
}
