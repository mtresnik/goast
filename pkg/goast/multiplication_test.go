package goast

import (
	"fmt"
	"testing"
)

func TestMultiplication_ToString(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var constant2 = Constant{Representation: 0.6 + 0.3i}
	var constant3 = Constant{Representation: 0.2}
	var constant4 = Constant{Representation: 0.3i}
	var multiplication1 = Multiplication{Values: []Operation{constant1, constant2}}
	var multiplication2 = Multiplication{Values: []Operation{constant3, constant4}}
	var multiplication3 = Multiplication{Values: []Operation{multiplication1, multiplication2}}
	fmt.Println(multiplication3.String())
}

func TestMultiplication_ToString2(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var variable2 = Variable{Name: "x"}
	var variable3 = Variable{Name: "y"}
	var variable4 = Variable{Name: "z"}
	var multiplication1 = Multiplication{Values: []Operation{variable1, variable2}}
	var multiplication2 = Multiplication{Values: []Operation{variable3, variable4}}
	var multiplication3 = Multiplication{Values: []Operation{multiplication1, multiplication2}}
	fmt.Println(multiplication3.String())
}

func TestMultiplication_DeepFlatten(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var variable2 = Variable{Name: "x"}
	var variable3 = Variable{Name: "y"}
	var variable4 = Variable{Name: "z"}
	var multiplication1 = Multiplication{Values: []Operation{variable1, variable2}}
	var multiplication2 = Multiplication{Values: []Operation{variable3, variable4}}
	var multiplication3 = Multiplication{Values: []Operation{multiplication1, multiplication2}}
	var flattened = OperationDeepFlatten(multiplication3)
	fmt.Println(OperationsToString(flattened))
}
