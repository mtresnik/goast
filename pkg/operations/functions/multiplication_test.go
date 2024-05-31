package functions

import (
	"fmt"
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/variables"
	"testing"
)

func TestMultiplication_ToString(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var constant2 = constants.Constant{Representation: 0.6 + 0.3i}
	var constant3 = constants.Constant{Representation: 0.2}
	var constant4 = constants.Constant{Representation: 0.3i}
	var multiplication1 = Multiplication{Values: []operations.Operation{constant1, constant2}}
	var multiplication2 = Multiplication{Values: []operations.Operation{constant3, constant4}}
	var multiplication3 = Multiplication{Values: []operations.Operation{multiplication1, multiplication2}}
	fmt.Println(multiplication3.ToString())
}

func TestMultiplication_ToString2(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var variable2 = variables.Variable{Name: "x"}
	var variable3 = variables.Variable{Name: "y"}
	var variable4 = variables.Variable{Name: "z"}
	var multiplication1 = Multiplication{Values: []operations.Operation{variable1, variable2}}
	var multiplication2 = Multiplication{Values: []operations.Operation{variable3, variable4}}
	var multiplication3 = Multiplication{Values: []operations.Operation{multiplication1, multiplication2}}
	fmt.Println(multiplication3.ToString())
}

func TestMultiplication_DeepFlatten(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var variable2 = variables.Variable{Name: "x"}
	var variable3 = variables.Variable{Name: "y"}
	var variable4 = variables.Variable{Name: "z"}
	var multiplication1 = Multiplication{Values: []operations.Operation{variable1, variable2}}
	var multiplication2 = Multiplication{Values: []operations.Operation{variable3, variable4}}
	var multiplication3 = Multiplication{Values: []operations.Operation{multiplication1, multiplication2}}
	var flattened = operations.DeepFlatten(multiplication3)
	fmt.Println(operations.ToString(flattened))
}
