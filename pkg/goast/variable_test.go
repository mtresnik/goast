package goast

import (
	"fmt"
	"testing"
)

func TestVariable_IsConstant(t *testing.T) {
	var variable1 = Variable{Name: "x"}
	var isConstant = AllConstants(variable1)
	fmt.Println(isConstant)
}
