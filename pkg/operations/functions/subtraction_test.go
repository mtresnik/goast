package functions

import (
	"fmt"
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/variables"
	"testing"
)

func TestSubtraction_ToString(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var constant1 = constants.Constant{Representation: 0.2}
	var subtraction1 = Subtraction{Values: []operations.Operation{variable1, constant1}}
	fmt.Println(subtraction1.ToString())
}
