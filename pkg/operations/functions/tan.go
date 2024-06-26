package functions

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
	"math/cmplx"
)

type Tan struct {
	Inner operations.Operation
}

func (t Tan) GetValues() []operations.Operation {
	return []operations.Operation{t.Inner}
}

func (t Tan) IsConstant() bool {
	return t.Inner.IsConstant()
}

func (t Tan) ToString() string {
	if t.Inner.IsConstant() {
		c := t.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := "tan("
	retString += t.Inner.ToString()
	retString += ")"
	return retString
}

func (t Tan) ToNumber() complex128 {
	if t.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Tan(t.Inner.ToNumber())
}

func (t Tan) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Tan{Inner: t.Inner.Evaluate(one, other)}
}
