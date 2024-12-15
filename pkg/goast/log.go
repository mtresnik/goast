package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Log struct {
	Base  Operation
	Inner Operation
}

func (l Log) GetValues() []Operation {
	return []Operation{l.Base, l.Inner}
}

func (l Log) IsConstant() bool {
	return l.Base.IsConstant() && l.Inner.IsConstant()
}

func (l Log) String() string {
	if l.Inner.IsConstant() {
		c := l.Number()
		return goutils.SmartComplexString(c)
	}
	retString := "log_("
	retString += l.Base.String() + ", "
	retString += l.Inner.String()
	retString += ")"
	return retString
}

func (l Log) Number() complex128 {
	if l.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Log(l.Inner.Number()) / cmplx.Log(l.Base.Number())
}

func (l Log) Evaluate(one Operation, other Operation) Operation {
	return Log{
		Base:  l.Base.Evaluate(one, other),
		Inner: l.Inner.Evaluate(one, other),
	}
}
