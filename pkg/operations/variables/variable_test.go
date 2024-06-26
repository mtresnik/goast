package variables

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/operations"
	"testing"
)

func TestVariable_IsConstant(t *testing.T) {
	var variable1 = Variable{Name: "x"}
	var isConstant = operations.AllConstants(variable1)
	fmt.Println(isConstant)
}
