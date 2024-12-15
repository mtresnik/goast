package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Addition struct {
	Values []Operation
}

func (a Addition) GetValues() []Operation {
	return a.Values
}

func (a Addition) IsConstant() bool {
	if len(a.Values) == 0 {
		return true
	}
	if len(a.Values) == 1 {
		return a.Values[0].IsConstant()
	}
	return AllConstants(a)
}

func (a Addition) String() string {
	if a.IsConstant() {
		c := a.Number()
		return goutils.SmartComplexString(c)
	}
	retString := ""
	var values = a.GetValues()
	for i := 0; i < len(values); i++ {
		retString += values[i].String()
		if i < len(values)-1 {
			retString += " + "
		}
	}
	return retString
}

func (a Addition) Number() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := 0 + 0i
	for i := 0; i < len(a.Values); i++ {
		ret += a.Values[i].Number()
	}
	return ret
}

func (a Addition) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(a, one) {
		return other
	}
	return Addition{Values: EvaluateValues(a, one, other)}
}
