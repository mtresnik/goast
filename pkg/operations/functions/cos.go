package functions

import (
	"goast/pkg/operations"
	"math/cmplx"
	"strconv"
)

type Cos struct {
	Inner operations.Operation
}

func (c Cos) GetValues() []operations.Operation {
	return []operations.Operation{c.Inner}
}

func (c Cos) IsConstant() bool {
	return c.Inner.IsConstant()
}

func (c Cos) ToString() string {
	if c.Inner.IsConstant() {
		c := c.ToNumber()
		return strconv.FormatComplex(c, 'f', 5, 64)
	}
	retString := "cos("
	retString += c.Inner.ToString()
	retString += ")"
	return retString
}

func (c Cos) ToNumber() complex128 {
	if c.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Cos(c.Inner.ToNumber())
}

func (c Cos) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Cos{
		Inner: c.Inner.Evaluate(one, other),
	}
}
