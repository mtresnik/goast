package constants

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations"
	"testing"
)

func TestConstant_IsConstant(t *testing.T) {
	var constant1 = Constant{Representation: 0.5 + 0.5i}
	var isConstant = operations.AllConstants(constant1)
	fmt.Println(isConstant)
}
