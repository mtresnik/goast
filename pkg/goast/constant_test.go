package goast

import (
	"fmt"
	"testing"
)

func TestConstant_IsConstant(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var isConstant = AllConstants(constant1)
	fmt.Println(isConstant)
}
