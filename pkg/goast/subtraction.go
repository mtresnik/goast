package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Subtraction struct {
	Values []Operation
}

func (s Subtraction) GetValues() []Operation {
	return s.Values
}

func (s Subtraction) IsConstant() bool {
	if len(s.Values) == 0 {
		return true
	}
	if len(s.Values) == 1 {
		return s.Values[0].IsConstant()
	}
	return AllConstants(s)
}

func (s Subtraction) String() string {
	if s.IsConstant() {
		c := s.Number()
		return goutils.SmartComplexString(c)
	}
	retString := ""
	for i := 0; i < len(s.Values); i++ {
		retString += s.Values[i].String()
		if i < len(s.Values)-1 {
			retString += " - "
		}
	}
	return retString
}

func (s Subtraction) Number() complex128 {
	if s.IsConstant() == false {
		return cmplx.NaN()
	}
	if len(s.Values) == 0 {
		return 0
	}
	retValue := s.Values[0].Number()
	for i := 1; i < len(s.Values); i++ {
		retValue -= s.Values[i].Number()
	}
	return retValue
}

func (s Subtraction) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(s, one) {
		return other
	}
	return Subtraction{Values: EvaluateValues(s, one, other)}
}
