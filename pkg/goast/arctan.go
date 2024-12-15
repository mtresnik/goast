package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type ArcTan struct {
	Inner Operation
}

func (a ArcTan) GetValues() []Operation {
	return []Operation{a.Inner}
}

func (a ArcTan) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a ArcTan) String() string {
	if a.Inner.IsConstant() {
		c := a.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "arctan("
	retString += a.Inner.String()
	retString += ")"
	return retString
}

func (a ArcTan) Number() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Atan(a.Inner.Number())
}

func (a ArcTan) Evaluate(one Operation, other Operation) Operation {
	return ArcTan{Inner: a.Inner.Evaluate(one, other)}
}
