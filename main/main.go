package main

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/goast"
	"github.com/mtresnik/goast/pkg/operations/functions"
)

func main() {
	TestParse()
}

func TestParse() {
	var inputString = "a * b * c + 123 + sin(x)"
	var parsed, err = goast.ParseOperation(inputString)
	if err != nil {
		fmt.Println((*err).Error())
		return
	}
	fmt.Println(parsed.String())
	parsed = parsed.Evaluate(goast.X, functions.Division{Numerator: goast.PI, Denominator: goast.TWO})
	fmt.Println(parsed.String())
}
