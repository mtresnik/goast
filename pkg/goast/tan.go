package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Tan struct {
	Inner Operation
}

func (t Tan) GetValues() []Operation {
	return []Operation{t.Inner}
}

func (t Tan) IsConstant() bool {
	return t.Inner.IsConstant()
}

func (t Tan) String() string {
	if t.Inner.IsConstant() {
		c := t.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "tan("
	retString += t.Inner.String()
	retString += ")"
	return retString
}

func (t Tan) Number() complex128 {
	if t.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Tan(t.Inner.Number())
}

func (t Tan) Evaluate(one Operation, other Operation) Operation {
	return Tan{Inner: t.Inner.Evaluate(one, other)}
}
