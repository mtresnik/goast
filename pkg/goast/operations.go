package goast

import (
	"reflect"
)

type Operation interface {
	GetValues() []Operation
	IsConstant() bool
	String() string
	Number() complex128
	Evaluate(one Operation, other Operation) Operation
}

func AllConstants(root Operation) bool {
	var flattened = OperationDeepFlatten(root)
	for i := 0; i < len(flattened); i++ {
		if flattened[i].IsConstant() == false {
			return false
		}
	}
	return true
}

func hasNestedValues(root []Operation) bool {
	for i := 0; i < len(root); i++ {
		var elem = root[i]
		if len(elem.GetValues()) > 1 {
			return true
		}
	}
	return false
}

func FlattenRoot(root Operation) []Operation {
	values := root.GetValues()
	if len(values) == 0 {
		return []Operation{root}
	}
	return OperationFlatten(values)
}

func OperationFlatten(values []Operation) []Operation {
	var count = 0
	if len(values) == 0 {
		return values
	}
	for i := 0; i < len(values); i++ {
		var elem = values[i]
		var childValues = elem.GetValues()
		if len(childValues) > 0 {
			count += len(childValues)
		} else {
			count++
		}
	}
	var retArray = make([]Operation, count)
	var outerIndex = 0
	for i := 0; i < len(values); i++ {
		var elem = values[i]
		var childValues = elem.GetValues()
		if len(childValues) > 0 {
			for j := 0; j < len(childValues); j++ {
				retArray[outerIndex] = childValues[j]
				outerIndex++
			}
		} else {
			retArray[outerIndex] = values[i]
			outerIndex++
		}
	}
	return retArray
}

func OperationDeepFlatten(root Operation) []Operation {
	retArray := []Operation{root}
	for hasNestedValues(retArray) {
		retArray = OperationFlatten(retArray)
	}
	return retArray
}

func OperationsToString(values []Operation) string {
	retString := "["
	for i := 0; i < len(values); i++ {
		var elem = values[i]
		retString += elem.String()
		if i < len(values)-1 {
			retString += ", "
		}
	}
	retString += "]"
	return retString
}

func FlatString(root Operation) string {
	return OperationsToString(FlattenRoot(root))
}

func Contains(op Operation, arr []Operation) bool {
	for i := 0; i < len(arr); i++ {
		if OperationEquals(op, arr[i]) {
			return true
		}
	}
	return false
}

func ContainsAll(arr1 []Operation, arr2 []Operation) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if Contains(arr1[i], arr2) == false {
			return false
		}
	}
	return true
}

func OperationEquals(one Operation, two Operation) bool {
	oneIsConstant := one.IsConstant()
	twoIsConstant := two.IsConstant()
	if oneIsConstant && twoIsConstant {
		return one.Number() == two.Number()
	}
	if oneIsConstant != twoIsConstant {
		return false
	}
	oneValues := one.GetValues()
	twoValues := two.GetValues()
	if len(oneValues) == 0 && len(twoValues) == 0 {
		return one.String() == two.String()
	}
	oneType := reflect.TypeOf(one)
	twoType := reflect.TypeOf(two)
	oneFlattened := OperationDeepFlatten(one)
	twoFlattened := OperationDeepFlatten(two)
	if oneType == twoType && len(oneFlattened) != len(twoFlattened) {
		return false
	}
	if oneType == twoType && ContainsAll(oneValues, twoValues) {
		return true
	}
	return one.String() == two.String()
}

func EvaluateValues(a Operation, one Operation, other Operation) []Operation {
	values := a.GetValues()
	retArray := make([]Operation, len(values))
	for i := 0; i < len(values); i++ {
		if OperationEquals(values[i], one) {
			retArray[i] = other
		} else {
			retArray[i] = values[i].Evaluate(one, other)
		}
	}
	return retArray
}
