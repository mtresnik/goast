package constants

import (
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/utils"
)

type Constant struct {
	Representation       complex128
	StringRepresentation *string
}

func (c Constant) GetValues() []operations.Operation {
	return make([]operations.Operation, 0)
}

func (c Constant) IsConstant() bool {
	return true
}

func (c Constant) ToString() string {
	if c.StringRepresentation != nil {
		return *(c.StringRepresentation)
	}
	return utils.SmartComplexString(c.Representation)
}

func (c Constant) ToNumber() complex128 {
	return c.Representation
}

func (c Constant) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(c, one) {
		return other
	}
	return c
}
