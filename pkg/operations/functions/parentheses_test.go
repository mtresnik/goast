package functions

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations"
	"github.com/mtresnik/goast/pkg/operations/constants"
	"github.com/mtresnik/goast/pkg/operations/variables"
	"testing"
)

func TestParentheses_ToString(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var variable1 = variables.Variable{Name: "x"}
	var addition1 = Addition{Values: []operations.Operation{constant1, variable1}}
	var parentheses1 = Parentheses{Inner: addition1}
	fmt.Println(parentheses1.ToString())
}

func TestParentheses_ToString2(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var variable1 = variables.Variable{Name: "x"}
	var addition1 = Addition{Values: []operations.Operation{constant1, variable1}}
	var parentheses1 = Parentheses{Inner: addition1}
	var parentheses2 = Parentheses{Inner: parentheses1}
	fmt.Println(parentheses2.ToString())
}
