package goast

import (
	"strconv"
)

var ReservedVariables = []string{eString, iString}

var E_Variable = Variable{Name: eString}
var I_Variable = Variable{Name: iString}
var X_Variable = Variable{Name: "x"}

func NewVariable(name string) Operation {
	return Variable{Name: name}
}

func GenerateVariables(size int) []Operation {
	retList := make([]Operation, 0)
	count := 0
	for i := 0; i < size; i++ {
		name := "x" + strconv.Itoa(i)
		retList = append(retList, Variable{Name: name})
		count++
	}
	return retList
}
