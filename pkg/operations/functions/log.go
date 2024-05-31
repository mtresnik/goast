package functions

import (
	"goast/pkg/operations"
	"math/cmplx"
	"strconv"
)

type Log struct {
	Base  operations.Operation
	Inner operations.Operation
}

func (l Log) GetValues() []operations.Operation {
	return []operations.Operation{l.Base, l.Inner}
}

func (l Log) IsConstant() bool {
	return l.Base.IsConstant() && l.Inner.IsConstant()
}

func (l Log) ToString() string {
	if l.Inner.IsConstant() {
		c := l.ToNumber()
		return strconv.FormatComplex(c, 'f', 5, 64)
	}
	retString := "log("
	retString += l.Base.ToString() + ", "
	retString += l.Inner.ToString()
	retString += ")"
	return retString
}

func (l Log) ToNumber() complex128 {
	if l.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Log(l.Inner.ToNumber()) / cmplx.Log(l.Base.ToNumber())
}

func (l Log) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Log{
		Base:  l.Base.Evaluate(one, other),
		Inner: l.Inner.Evaluate(one, other),
	}
}
