package goast

import (
	"fmt"
	"testing"
)

func TestSubtraction_ToString(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var constant1 = Constant{Representation: 0.2}
	var subtraction1 = Subtraction{Values: []Operation{variable1, constant1}}
	fmt.Println(subtraction1.String())
}
