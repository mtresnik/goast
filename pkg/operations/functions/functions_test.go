package functions

import (
	"fmt"
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/variables"
	"testing"
)

func TestAddFunction(t *testing.T) {
	builderFunction := func(ops ...operations.Operation) operations.Operation {
		var value1 = ops[0]
		var value2 = ops[1]
		var pow1 = Power{Base: value1, Exponent: constants.TWO}
		var addition1 = Addition{Values: []operations.Operation{pow1, value2}}
		return addition1
	}
	var builder = FunctionBuilder{
		NumberOfParams: 2,
		Function:       builderFunction,
	}
	AddFunction("f", builder)
}

func TestBuildFunction(t *testing.T) {
	var function = BuildFunction("sin", variables.Variable{Name: "x"})
	fmt.Println(function.ToString())
	var evaluated = function.Evaluate(variables.Variable{Name: "x"}, constants.PI)
	fmt.Println(evaluated.ToString())
}

func TestBuildFunction2(t *testing.T) {
	var function = BuildFunction("abc", variables.Variable{Name: "x"})
	fmt.Println(function.ToString())
}
