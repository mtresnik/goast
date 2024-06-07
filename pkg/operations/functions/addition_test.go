package functions

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/operations/constants"
	"github.com/mtresnik/goast/pkg/operations/variables"
	"testing"
)

func TestAddition_GetValues(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var variable1 = variables.Variable{Name: "x"}
	var addition1 = Addition{Values: []operations.Operation{constant1, variable1}}
	fmt.Println(addition1.ToString())
}

func TestAddition_IsConstant(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var variable1 = variables.Variable{Name: "x"}
	var addition1 = Addition{Values: []operations.Operation{constant1, variable1}}
	fmt.Println(addition1.IsConstant())
}

func TestAddition_ToNumber(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var variable1 = variables.Variable{Name: "x"}
	var addition1 = Addition{Values: []operations.Operation{constant1, variable1}}
	fmt.Println(addition1.ToNumber())
}

func TestAddition_ToNumber2(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var constant2 = constants.Constant{Representation: 0.6 + 0.3i}
	var addition1 = Addition{Values: []operations.Operation{constant1, constant2}}
	fmt.Println(addition1.ToNumber())
}

func TestAddition_Nested(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var constant2 = constants.Constant{Representation: 0.6 + 0.3i}
	var constant3 = constants.Constant{Representation: 0.2}
	var constant4 = constants.Constant{Representation: 0.3i}
	var addition1 = Addition{Values: []operations.Operation{constant1, constant2}}
	var addition2 = Addition{Values: []operations.Operation{constant3, constant4}}
	var addition3 = Addition{Values: []operations.Operation{addition1, addition2}}
	fmt.Println(addition3.ToString())
}

func TestAddition_Nested2(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var variable2 = variables.Variable{Name: "x"}
	var variable3 = variables.Variable{Name: "y"}
	var variable4 = variables.Variable{Name: "z"}
	var addition1 = Addition{Values: []operations.Operation{variable1, variable2}}
	var addition2 = Addition{Values: []operations.Operation{variable3, variable4}}
	var addition3 = Addition{Values: []operations.Operation{addition1, addition2}}
	fmt.Println(addition3.ToString())
}

func TestAddition_ToString(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var variable2 = variables.Variable{Name: "x"}
	var variable3 = variables.Variable{Name: "y"}
	var variable4 = variables.Variable{Name: "z"}
	var addition1 = Addition{Values: []operations.Operation{variable1, variable2}}
	var addition2 = Addition{Values: []operations.Operation{variable3, variable4}}
	var addition3 = Addition{Values: []operations.Operation{addition1, addition2}}
	fmt.Println(operations.ToString(addition3.Values))
}

func TestAddition_ToString2(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var variable2 = variables.Variable{Name: "x"}
	var variable3 = variables.Variable{Name: "y"}
	var variable4 = variables.Variable{Name: "z"}
	var addition1 = Addition{Values: []operations.Operation{variable1, variable2}}
	var addition2 = Addition{Values: []operations.Operation{variable3, variable4}}
	var addition3 = Addition{Values: []operations.Operation{addition1, addition2}}
	fmt.Println(operations.FlatString(addition3))
}

func TestAddition_Evaluate(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var constant1 = constants.Constant{Representation: 0.2}
	var addition1 = Addition{Values: []operations.Operation{variable1, constant1}}
	fmt.Println(addition1.ToString())
	var addition2 = addition1.Evaluate(variable1, constants.Constant{Representation: 0.8})
	fmt.Println(addition2.ToString())
}

func TestAddition_Equals(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var constant1 = constants.Constant{Representation: 1}
	var addition1 = Addition{Values: []operations.Operation{variable1, constant1}}
	var addition2 = Addition{Values: []operations.Operation{constant1, variable1}}
	fmt.Println(operations.Equals(addition1, addition2))
}
