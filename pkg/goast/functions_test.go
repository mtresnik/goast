package goast

import (
	"fmt"
	"testing"
)

func TestAddFunction(t *testing.T) {
	builderFunction := func(ops ...Operation) Operation {
		var value1 = ops[0]
		var value2 = ops[1]
		var pow1 = Power{Base: value1, Exponent: TWO_Constant}
		var addition1 = Addition{Values: []Operation{pow1, value2}}
		return addition1
	}
	var builder = FunctionBuilder{
		NumberOfParams: 2,
		Function:       builderFunction,
	}
	AddFunction("f", builder)
}

func TestBuildFunction(t *testing.T) {
	var function = BuildFunction("sin", Variable{Name: "x"})
	fmt.Println(function.String())
	var evaluated = function.Evaluate(Variable{Name: "x"}, PI_Constant)
	fmt.Println(evaluated.String())
}

func TestBuildFunction2(t *testing.T) {
	var function = BuildFunction("abc", Variable{Name: "x"})
	fmt.Println(function.String())
}
