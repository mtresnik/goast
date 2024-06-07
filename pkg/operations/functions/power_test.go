package functions

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations/constants"
	"github.com/mtresnik/goast/pkg/operations/variables"
	"testing"
)

func TestPower_ToString(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var variable2 = variables.Variable{Name: "x"}
	var power1 = Power{Base: variable1, Exponent: variable2}
	fmt.Println(power1.ToString())
}

func TestPower_ToString2(t *testing.T) {
	var variable1 = variables.Variable{Name: "w"}
	var constant1 = constants.Constant{Representation: 0}
	var power1 = Power{Base: variable1, Exponent: constant1}
	fmt.Println(power1.ToString())
}
