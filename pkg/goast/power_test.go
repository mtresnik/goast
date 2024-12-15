package goast

import (
	"fmt"
	"testing"
)

func TestPower_ToString(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var variable2 = Variable{Name: "x"}
	var power1 = Power{Base: variable1, Exponent: variable2}
	fmt.Println(power1.String())
}

func TestPower_ToString2(t *testing.T) {
	var variable1 = Variable{Name: "w"}
	var constant1 = Constant{Representation: 0}
	var power1 = Power{Base: variable1, Exponent: constant1}
	fmt.Println(power1.String())
}
