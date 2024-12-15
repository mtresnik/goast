package goast

import (
	"fmt"
	"testing"
)

func TestParentheses_ToString(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var variable1 = Variable{Name: "x"}
	var addition1 = Addition{Values: []Operation{constant1, variable1}}
	var parentheses1 = Parentheses{Inner: addition1}
	fmt.Println(parentheses1.String())
}

func TestParentheses_ToString2(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var variable1 = Variable{Name: "x"}
	var addition1 = Addition{Values: []Operation{constant1, variable1}}
	var parentheses1 = Parentheses{Inner: addition1}
	var parentheses2 = Parentheses{Inner: parentheses1}
	fmt.Println(parentheses2.String())
}
