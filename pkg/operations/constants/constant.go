package constants

import (
	"goast/pkg/operations"
	"strconv"
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
	return strconv.FormatComplex(c.Representation, 'G', 5, 64)
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
