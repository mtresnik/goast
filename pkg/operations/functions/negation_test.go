package functions

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations/constants"
	"testing"
)

func TestNegation_ToString(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var negation1 = Negation{Inner: constant1}
	fmt.Println(negation1.ToString())
}

func TestNegation_ToNumber(t *testing.T) {
	var constant1 = constants.Constant{Representation: 0.5 + 0.5i}
	var negation1 = Negation{Inner: constant1}
	fmt.Println(negation1.ToNumber())
}
