package functions

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
	"math/cmplx"
)

type ArcTan struct {
	Inner operations.Operation
}

func (a ArcTan) GetValues() []operations.Operation {
	return []operations.Operation{a.Inner}
}

func (a ArcTan) IsConstant() bool {
	return a.Inner.IsConstant()
}

func (a ArcTan) ToString() string {
	if a.Inner.IsConstant() {
		c := a.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := "arctan("
	retString += a.Inner.ToString()
	retString += ")"
	return retString
}

func (a ArcTan) ToNumber() complex128 {
	if a.IsConstant() == false {
		return cmplx.NaN()
	}
	return cmplx.Atan(a.Inner.ToNumber())
}

func (a ArcTan) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return ArcTan{Inner: a.Inner.Evaluate(one, other)}
}
