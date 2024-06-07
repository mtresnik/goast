package functions

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
	"math/cmplx"
)

type Addition struct {
	Values []operations.Operation
}

func (a Addition) GetValues() []operations.Operation {
	return a.Values
}

func (a Addition) IsConstant() bool {
	if len(a.Values) == 0 {
		return true
	}
	if len(a.Values) == 1 {
		return a.Values[0].IsConstant()
	}
	return operations.AllConstants(a)
}

func (a Addition) ToString() string {
	if a.IsConstant() {
		c := a.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := ""
	var values = a.GetValues()
	for i := 0; i < len(values); i++ {
		retString += values[i].ToString()
		if i < len(values)-1 {
			retString += " + "
		}
	}
	return retString
}

func (a Addition) ToNumber() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := 0 + 0i
	for i := 0; i < len(a.Values); i++ {
		ret += a.Values[i].ToNumber()
	}
	return ret
}

func (a Addition) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(a, one) {
		return other
	}
	return Addition{Values: operations.EvaluateValues(a, one, other)}
}
