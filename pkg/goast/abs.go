package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Abs struct {
	Inner Operation
}

func (a Abs) GetValues() []Operation {
	return []Operation{a.Inner}
}

func (a Abs) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a Abs) String() string {
	if a.Inner.IsConstant() {
		c := a.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "abs("
	retString += a.Inner.String()
	retString += ")"
	return retString
}

func (a Abs) Number() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return complex(cmplx.Abs(a.Inner.Number()), 0)
}

func (a Abs) Evaluate(one Operation, other Operation) Operation {
	return Abs{
		Inner: a.Inner.Evaluate(one, other),
	}
}
