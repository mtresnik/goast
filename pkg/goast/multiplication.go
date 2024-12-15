package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Multiplication struct {
	Values []Operation
}

func (m Multiplication) GetValues() []Operation {
	return m.Values
}

func (m Multiplication) IsConstant() bool {
	if len(m.Values) == 0 {
		return true
	}
	if len(m.Values) == 1 {
		return m.Values[0].IsConstant()
	}
	return AllConstants(m)
}

func (m Multiplication) String() string {
	if m.IsConstant() {
		c := m.Number()
		return goutils.SmartComplexString(c)
	}
	retString := ""
	var values = m.GetValues()
	for i := 0; i < len(values); i++ {
		retString += values[i].String()
		if i < len(values)-1 {
			retString += " * "
		}
	}
	return retString
}

func (m Multiplication) Number() complex128 {
	if m.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := 1 + 0i
	for i := 0; i < len(m.Values); i++ {
		ret *= m.Values[i].Number()
	}
	return ret
}

func (m Multiplication) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(m, one) {
		return other
	}
	return Multiplication{Values: EvaluateValues(m, one, other)}
}
