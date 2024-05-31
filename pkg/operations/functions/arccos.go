package functions

import (
	"goast/pkg/operations"
	"math/cmplx"
	"strconv"
)

type ArcCos struct {
	Inner operations.Operation
}

func (a ArcCos) GetValues() []operations.Operation {
	return []operations.Operation{a.Inner}
}

func (a ArcCos) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a ArcCos) ToString() string {
	if a.Inner.IsConstant() {
		c := a.ToNumber()
		return strconv.FormatComplex(c, 'f', 5, 64)
	}
	retString := "arccos("
	retString += a.Inner.ToString()
	retString += ")"
	return retString
}

func (a ArcCos) ToNumber() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Acos(a.Inner.ToNumber())
}

func (a ArcCos) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return ArcCos{Inner: a.Inner.Evaluate(one, other)}
}
