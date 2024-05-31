package parser

import (
	"fmt"
	"goast/pkg/operations/parser/tokens"
	"testing"
)

func TestNumber_Compile(t *testing.T) {
	var stringRep = "5+  2i"
	var tokenRep = tokens.NullIndex(tokens.Number, &stringRep)
	var number = INumber{
		StartIndex: 0,
		EndIndex:   1,
		Token:      tokenRep,
	}
	fmt.Println(number.ToString())
	var compiled = number.Compile()
	fmt.Println(compiled.ToString())
}

func TestAddition_Compile(t *testing.T) {
	var stringRep1 = "5+  2i"
	var tokenRep1 = tokens.NullIndex(tokens.Number, &stringRep1)
	var number1 = INumber{
		StartIndex: 0,
		EndIndex:   1,
		Token:      tokenRep1,
	}

	var stringRep2 = "6+  3i"
	var tokenRep2 = tokens.NullIndex(tokens.Number, &stringRep2)
	var number2 = INumber{
		StartIndex: 2,
		EndIndex:   3,
		Token:      tokenRep2,
	}

	var addition = IAddition{
		StartIndex: 0,
		EndIndex:   1,
		Left:       number1,
		Right:      number2,
	}
	fmt.Println(addition.ToString())
	var compiled = addition.Compile()
	fmt.Println(compiled.ToString())
}
