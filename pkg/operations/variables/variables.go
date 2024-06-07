package variables

import (
	"github.com/mtresnik/goast/pkg/operations"
	"strconv"
)

var eString = "e"
var iString = "i"

var Reserved = []string{eString, iString}

var E = Variable{Name: eString}
var I = Variable{Name: iString}
var X = Variable{Name: "x"}

func New(name string) operations.Operation {
	return Variable{Name: name}
}

func GenerateVariables(size int) []operations.Operation {
	retList := make([]operations.Operation, 0)
	count := 0
	for i := 0; i < size; i++ {
		name := "x" + strconv.Itoa(i)
		retList = append(retList, Variable{Name: name})
		count++
	}
	return retList
}
