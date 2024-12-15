package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Cos struct {
	Inner Operation
}

func (c Cos) GetValues() []Operation {
	return []Operation{c.Inner}
}

func (c Cos) IsConstant() bool {
	return c.Inner.IsConstant()
}

func (c Cos) String() string {
	if c.Inner.IsConstant() {
		c := c.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "cos("
	retString += c.Inner.String()
	retString += ")"
	return retString
}

func (c Cos) Number() complex128 {
	if c.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Cos(c.Inner.Number())
}

func (c Cos) Evaluate(one Operation, other Operation) Operation {
	return Cos{
		Inner: c.Inner.Evaluate(one, other),
	}
}
