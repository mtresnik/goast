package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Parentheses struct {
	Inner Operation
}

func (p Parentheses) GetValues() []Operation {
	return []Operation{p.Inner}
}

func (p Parentheses) IsConstant() bool {
	return p.Inner.IsConstant()
}

func (p Parentheses) String() string {
	if p.Inner.IsConstant() {
		c := p.Inner.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "("
	retString += p.Inner.String()
	retString += ")"
	return retString
}

func (p Parentheses) Number() complex128 {
	if p.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := p.Inner.Number()
	return ret
}

func (p Parentheses) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(p, one) {
		return other
	}
	return Parentheses{Inner: p.Inner.Evaluate(one, other)}
}
