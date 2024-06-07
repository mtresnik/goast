package functions

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
	"math/cmplx"
)

type Negation struct {
	Inner operations.Operation
}

func (n Negation) GetValues() []operations.Operation {
	return []operations.Operation{n.Inner}
}

func (n Negation) IsConstant() bool {
	return n.Inner.IsConstant()
}

func (n Negation) ToString() string {
	if n.Inner.IsConstant() {
		c := n.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := "-"
	retString += n.Inner.ToString()
	return retString
}

func (n Negation) ToNumber() complex128 {
	if n.IsConstant() == false {
		return cmplx.NaN()
	}
	ret := 0 + 0i
	ret -= n.Inner.ToNumber()
	return ret
}

func (n Negation) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(n, one) {
		return other
	}
	return Negation{Inner: n.Inner.Evaluate(one, other)}
}
