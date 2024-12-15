package goast

import (
	"fmt"
	"testing"
)

func TestParse1(t *testing.T) {
	var inputString = "a + 123 + sin(x)"
	var parsed, err = ParseOperation(inputString)
	if err != nil {
		t.Error(*err)
	}
	fmt.Println(parsed.String())
	parsed = parsed.Evaluate(X_Variable, Division{Numerator: PI_Constant, Denominator: TWO_Constant})
	fmt.Println(parsed.String())
	parsed = parsed.Evaluate(Variable{Name: "a"}, TEN_Constant)
	fmt.Println(parsed.String())
}

func TestParse2(t *testing.T) {
	var inputString = "log_(x,y)"
	var parsed, err = ParseOperation(inputString)
	if err != nil {
		t.Error(*err)
	}
	fmt.Println(parsed.String())
}
