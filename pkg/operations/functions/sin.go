package functions

import (
	"goast/pkg/operations"
	"goast/pkg/utils"
	"math/cmplx"
)

type Sin struct {
	Inner operations.Operation
}

func (s Sin) GetValues() []operations.Operation {
	return []operations.Operation{s.Inner}
}

func (s Sin) IsConstant() bool {
	return s.Inner.IsConstant()
}

func (s Sin) ToString() string {
	if s.Inner.IsConstant() {
		c := s.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := "sin("
	retString += s.Inner.ToString()
	retString += ")"
	return retString
}

func (s Sin) ToNumber() complex128 {
	if s.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Sin(s.Inner.ToNumber())
}

func (s Sin) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Sin{
		Inner: s.Inner.Evaluate(one, other),
	}
}
