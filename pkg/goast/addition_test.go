package goast

import (
	"fmt"
	"testing"
)

func TestAddition_GetValues(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var variable1 = Variable{Name: "x"}
	var addition1 = Addition{Values: []Operation{constant1, variable1}}
	fmt.Println(addition1.String())
}

func TestAddition_IsConstant(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var variable1 = Variable{Name: "x"}
	var addition1 = Addition{Values: []Operation{constant1, variable1}}
	fmt.Println(addition1.IsConstant())
}

func TestAddition_ToNumber(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var variable1 = Variable{Name: "x"}
	var addition1 = Addition{Values: []Operation{constant1, variable1}}
	fmt.Println(addition1.Number())
}

func TestAddition_ToNumber2(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var constant2 = Constant{Representation: 0.6 + 0.3i}
	var addition1 = Addition{Values: []Operation{constant1, constant2}}
	fmt.Println(addition1.Number())
}

func TestAddition_Nested(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var constant2 = Constant{Representation: 0.6 + 0.3i}
	var constant3 = Constant{Representation: 0.2}
	var constant4 = Constant{Representation: 0.3i}
	var addition1 = Addition{Values: []Operation{constant1, constant2}}
	var addition2 = Addition{Values: []Operation{constant3, constant4}}
	var addition3 = Addition{Values: []Operation{addition1, addition2}}
	fmt.Println(addition3.String())
}

func TestAddition_Nested2(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var variable2 = Variable{Name: "x"}
	var variable3 = Variable{Name: "y"}
	var variable4 = Variable{Name: "z"}
	var addition1 = Addition{Values: []Operation{variable1, variable2}}
	var addition2 = Addition{Values: []Operation{variable3, variable4}}
	var addition3 = Addition{Values: []Operation{addition1, addition2}}
	fmt.Println(addition3.String())
}

func TestAddition_ToString(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var variable2 = Variable{Name: "x"}
	var variable3 = Variable{Name: "y"}
	var variable4 = Variable{Name: "z"}
	var addition1 = Addition{Values: []Operation{variable1, variable2}}
	var addition2 = Addition{Values: []Operation{variable3, variable4}}
	var addition3 = Addition{Values: []Operation{addition1, addition2}}
	fmt.Println(OperationsToString(addition3.Values))
}

func TestAddition_ToString2(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var variable2 = Variable{Name: "x"}
	var variable3 = Variable{Name: "y"}
	var variable4 = Variable{Name: "z"}
	var addition1 = Addition{Values: []Operation{variable1, variable2}}
	var addition2 = Addition{Values: []Operation{variable3, variable4}}
	var addition3 = Addition{Values: []Operation{addition1, addition2}}
	fmt.Println(FlatString(addition3))
}

func TestAddition_Evaluate(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var constant1 = Constant{Representation: 0.2}
	var addition1 = Addition{Values: []Operation{variable1, constant1}}
	fmt.Println(addition1.String())
	var addition2 = addition1.Evaluate(variable1, Constant{Representation: 0.8})
	fmt.Println(addition2.String())
}

func TestAddition_Equals(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var constant1 = Constant{Representation: 1}
	var addition1 = Addition{Values: []Operation{variable1, constant1}}
	var addition2 = Addition{Values: []Operation{constant1, variable1}}
	fmt.Println(OperationEquals(addition1, addition2))
}
