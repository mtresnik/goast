package goast

import "github.com/mtresnik/goutils/pkg/goutils"

type Constant struct {
	Representation       complex128
	StringRepresentation *string
}

func (c Constant) GetValues() []Operation {
	return make([]Operation, 0)
}

func (c Constant) IsConstant() bool {
	return true
}

func (c Constant) String() string {
	if c.StringRepresentation != nil {
		return *(c.StringRepresentation)
	}
	return goutils.SmartComplexString(c.Representation)
}

func (c Constant) Number() complex128 {
	return c.Representation
}

func (c Constant) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(c, one) {
		return other
	}
	return c
}
