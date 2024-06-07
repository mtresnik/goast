package functions

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
	"math/cmplx"
)

type Multiplication struct {
	Values []operations.Operation
}

func (m Multiplication) GetValues() []operations.Operation {
	return m.Values
}

func (m Multiplication) IsConstant() bool {
	if len(m.Values) == 0 {
		return true
	}
	if len(m.Values) == 1 {
		return m.Values[0].IsConstant()
	}
	return operations.AllConstants(m)
}

func (m Multiplication) ToString() string {
	if m.IsConstant() {
		c := m.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := ""
	var values = m.GetValues()
	for i := 0; i < len(values); i++ {
		retString += values[i].ToString()
		if i < len(values)-1 {
			retString += " * "
		}
	}
	return retString
}

func (m Multiplication) ToNumber() complex128 {
	if m.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := 1 + 0i
	for i := 0; i < len(m.Values); i++ {
		ret *= m.Values[i].ToNumber()
	}
	return ret
}

func (m Multiplication) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(m, one) {
		return other
	}
	return Multiplication{Values: operations.EvaluateValues(m, one, other)}
}
