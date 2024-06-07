package parser

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations/constants"
	"github.com/mtresnik/goast/pkg/operations/functions"
	"github.com/mtresnik/goast/pkg/operations/variables"
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
