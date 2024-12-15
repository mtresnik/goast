package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type ArcSin struct {
	Inner Operation
}

func (a ArcSin) GetValues() []Operation {
	return []Operation{a.Inner}
}

func (a ArcSin) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a ArcSin) String() string {
	if a.Inner.IsConstant() {
		c := a.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "arcsin("
	retString += a.Inner.String()
	retString += ")"
	return retString
}

func (a ArcSin) Number() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Asin(a.Inner.Number())
}

func (a ArcSin) Evaluate(one Operation, other Operation) Operation {
	return ArcSin{Inner: a.Inner.Evaluate(one, other)}
}
