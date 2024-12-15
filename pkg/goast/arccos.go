package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type ArcCos struct {
	Inner Operation
}

func (a ArcCos) GetValues() []Operation {
	return []Operation{a.Inner}
}

func (a ArcCos) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a ArcCos) String() string {
	if a.Inner.IsConstant() {
		c := a.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "arccos("
	retString += a.Inner.String()
	retString += ")"
	return retString
}

func (a ArcCos) Number() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Acos(a.Inner.Number())
}

func (a ArcCos) Evaluate(one Operation, other Operation) Operation {
	return ArcCos{Inner: a.Inner.Evaluate(one, other)}
}
