package goast

import "github.com/mtresnik/goutils/pkg/goutils"

var runtimeFunctions = map[string]FunctionBuilder{
	"sin":    {1, func(operations ...Operation) Operation { return Sin{Inner: operations[0]} }},
	"cos":    {1, func(operations ...Operation) Operation { return Cos{Inner: operations[0]} }},
	"tan":    {1, func(operations ...Operation) Operation { return Tan{Inner: operations[0]} }},
	"arcsin": {1, func(operations ...Operation) Operation { return ArcSin{Inner: operations[0]} }},
	"arccos": {1, func(operations ...Operation) Operation { return ArcCos{Inner: operations[0]} }},
	"arctan": {1, func(operations ...Operation) Operation { return ArcTan{Inner: operations[0]} }},
	"abs":    {1, func(operations ...Operation) Operation { return Abs{Inner: operations[0]} }},
	"log": {1, func(operations ...Operation) Operation {
		return Log{Base: TEN_Constant, Inner: operations[0]}
	}},
	"ln": {1, func(operations ...Operation) Operation {
		return Log{Base: E_Constant, Inner: operations[0]}
	}},
	"log_": {2, func(operations ...Operation) Operation {
		return Log{Base: operations[0], Inner: operations[1]}
	}},
}

var ReservedFunctions = goutils.Keys(runtimeFunctions)

// FunctionBuilder You must specify the number of params required for the building function.
type FunctionBuilder struct {
	NumberOfParams int
	Function       func(...Operation) Operation
}

func AddFunction(name string, builder FunctionBuilder) Operation {
	runtimeFunctions[name] = builder
	ReservedFunctions = goutils.Keys(runtimeFunctions)
	params := GenerateVariables(builder.NumberOfParams)
	return builder.Function(params...)
}

func BuildFunction(name string, params ...Operation) Operation {
	builder, exists := runtimeFunctions[name]
	if exists {
		if len(params) >= builder.NumberOfParams {
			return builder.Function(params...)
		}
	}
	values := make([]Operation, 0)
	for _, v := range name {
		var rep = string(v)
		values = append(values, Variable{
			Name: rep,
		})
	}
	return Multiplication{Values: values}
}
