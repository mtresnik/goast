package goast

import (
	"fmt"
	"testing"
)

func TestNegation_ToString(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var negation1 = Negation{Inner: constant1}
	fmt.Println(negation1.String())
}

func TestNegation_ToNumber(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var negation1 = Negation{Inner: constant1}
	fmt.Println(negation1.Number())
}
