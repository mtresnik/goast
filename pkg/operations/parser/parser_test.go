package parser

import (
	"fmt"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/functions"
	"goast/pkg/operations/variables"
	"testing"
)

func TestParse1(t *testing.T) {
	var inputString = "a + 123 + sin(x)"
	var parsed, err = ParseOperation(inputString)
	if err != nil {
		t.Error(*err)
	}
	fmt.Println(parsed.ToString())
	parsed = parsed.Evaluate(variables.X, functions.Division{Numerator: constants.PI, Denominator: constants.TWO})
	fmt.Println(parsed.ToString())
	parsed = parsed.Evaluate(variables.Variable{Name: "a"}, constants.TEN)
	fmt.Println(parsed.ToString())
}

func TestParse2(t *testing.T) {
	var inputString = "log_(x,y)"
	var parsed, err = ParseOperation(inputString)
	if err != nil {
		t.Error(*err)
	}
	fmt.Println(parsed.ToString())
}
