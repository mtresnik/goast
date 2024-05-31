package functions

import (
	"goast/pkg/operations"
	"math/cmplx"
	"strconv"
)

type Parentheses struct {
	Inner operations.Operation
}

func (p Parentheses) GetValues() []operations.Operation {
	return []operations.Operation{p.Inner}
}

func (p Parentheses) IsConstant() bool {
	return p.Inner.IsConstant()
}

func (p Parentheses) ToString() string {
	if p.Inner.IsConstant() {
		c := p.Inner.ToNumber()
		return strconv.FormatComplex(c, 'f', 5, 64)
	}
	retString := "("
	retString += p.Inner.ToString()
	retString += ")"
	return retString
}

func (p Parentheses) ToNumber() complex128 {
	if p.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := p.Inner.ToNumber()
	return ret
}

func (p Parentheses) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(p, one) {
		return other
	}
	return Parentheses{Inner: p.Inner.Evaluate(one, other)}
}
