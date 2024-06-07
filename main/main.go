package main

import (
	"fmt"
"github.com/mtresnik/goast/pkg/operations/constants"
"github.com/mtresnik/goast/pkg/operations/functions"
"github.com/mtresnik/goast/pkg/operations/parser"
"github.com/mtresnik/goast/pkg/operations/variables"
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
