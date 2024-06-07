package functions

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
	"math/cmplx"
)

type Subtraction struct {
	Values []operations.Operation
}

func (s Subtraction) GetValues() []operations.Operation {
	return s.Values
}

func (s Subtraction) IsConstant() bool {
	if len(s.Values) == 0 {
		return true
	}
	if len(s.Values) == 1 {
		return s.Values[0].IsConstant()
	}
	return operations.AllConstants(s)
}

func (s Subtraction) ToString() string {
	if s.IsConstant() {
		c := s.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := ""
	for i := 0; i < len(s.Values); i++ {
		retString += s.Values[i].ToString()
		if i < len(s.Values)-1 {
			retString += " - "
		}
	}
	return retString
}

func (s Subtraction) ToNumber() complex128 {
	if s.IsConstant() == false {
		return cmplx.NaN()
	}
	if len(s.Values) == 0 {
		return 0
	}
	retValue := s.Values[0].ToNumber()
	for i := 1; i < len(s.Values); i++ {
		retValue -= s.Values[i].ToNumber()
	}
	return retValue
}

func (s Subtraction) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(s, one) {
		return other
	}
	return Subtraction{Values: operations.EvaluateValues(s, one, other)}
}
