package functions

import (
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/variables"
	"goast/pkg/utils"
)

var runtimeFunctions = map[string]FunctionBuilder{
	"sin":    {1, func(operations ...operations.Operation) operations.Operation { return Sin{Inner: operations[0]} }},
	"cos":    {1, func(operations ...operations.Operation) operations.Operation { return Cos{Inner: operations[0]} }},
	"tan":    {1, func(operations ...operations.Operation) operations.Operation { return Tan{Inner: operations[0]} }},
	"arcsin": {1, func(operations ...operations.Operation) operations.Operation { return ArcSin{Inner: operations[0]} }},
	"arccos": {1, func(operations ...operations.Operation) operations.Operation { return ArcCos{Inner: operations[0]} }},
	"arctan": {1, func(operations ...operations.Operation) operations.Operation { return ArcTan{Inner: operations[0]} }},
	"abs":    {1, func(operations ...operations.Operation) operations.Operation { return Abs{Inner: operations[0]} }},
	"log": {1, func(operations ...operations.Operation) operations.Operation {
		return Log{Base: constants.TEN, Inner: operations[0]}
	}},
	"ln": {1, func(operations ...operations.Operation) operations.Operation {
		return Log{Base: constants.E, Inner: operations[0]}
	}},
	"log_": {2, func(operations ...operations.Operation) operations.Operation {
		return Log{Base: operations[0], Inner: operations[1]}
	}},
}

var Reserved = utils.Keys(runtimeFunctions)

// FunctionBuilder You must specify the number of params required for the building function.
type FunctionBuilder struct {
	NumberOfParams int
	Function       func(...operations.Operation) operations.Operation
}

func AddFunction(name string, builder FunctionBuilder) operations.Operation {
	runtimeFunctions[name] = builder
	Reserved = utils.Keys(runtimeFunctions)
	params := variables.GenerateVariables(builder.NumberOfParams)
	return builder.Function(params...)
}

func BuildFunction(name string, params ...operations.Operation) operations.Operation {
	builder, exists := runtimeFunctions[name]
	if exists {
		if len(params) >= builder.NumberOfParams {
			return builder.Function(params...)
		}
	}
	values := make([]operations.Operation, 0)
	for _, v := range name {
		var rep = string(v)
		values = append(values, variables.Variable{
			Name: rep,
		})
	}
	return Multiplication{Values: values}
}
