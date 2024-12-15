package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Negation struct {
	Inner Operation
}

func (n Negation) GetValues() []Operation {
	return []Operation{n.Inner}
}

func (n Negation) IsConstant() bool {
	return n.Inner.IsConstant()
}

func (n Negation) String() string {
	if n.Inner.IsConstant() {
		c := n.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "-"
	retString += n.Inner.String()
	return retString
}

func (n Negation) Number() complex128 {
	if n.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := 0 + 0i
	ret -= n.Inner.Number()
	return ret
}

func (n Negation) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(n, one) {
		return other
	}
	return Negation{Inner: n.Inner.Evaluate(one, other)}
}
