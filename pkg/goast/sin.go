package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Sin struct {
	Inner Operation
}

func (s Sin) GetValues() []Operation {
	return []Operation{s.Inner}
}

func (s Sin) IsConstant() bool {
	return s.Inner.IsConstant()
}

func (s Sin) String() string {
	if s.Inner.IsConstant() {
		c := s.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "sin("
	retString += s.Inner.String()
	retString += ")"
	return retString
}

func (s Sin) Number() complex128 {
	if s.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Sin(s.Inner.Number())
}

func (s Sin) Evaluate(one Operation, other Operation) Operation {
	return Sin{
		Inner: s.Inner.Evaluate(one, other),
	}
}
