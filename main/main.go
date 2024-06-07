package main

import (
	"fmt"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/functions"
	"goast/pkg/operations/parser"
	"goast/pkg/operations/variables"
)

func main() {
	TestParse()
}

func TestParse() {
	var inputString = "a * b * c + 123 + sin(x)"
	var parsed, err = parser.ParseOperation(inputString)
	if err != nil {
		fmt.Println((*err).Error())
		return
	}
	fmt.Println(parsed.ToString())
	parsed = parsed.Evaluate(variables.X, functions.Division{Numerator: constants.PI, Denominator: constants.TWO})
	fmt.Println(parsed.ToString())
}
