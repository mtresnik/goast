package utils

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestDeepFlatten(t *testing.T) {
	var tempArray = []any{[]any{[]any{0}}, 1, 2, 3, []any{[]any{[]any{4}}}, []any{5, 6, 7}}
	fmt.Println(tempArray)
	var flattened = deepFlatten(tempArray)
	fmt.Println(flattened)
}

func TestStringEndsWith(t *testing.T) {
	testString := "abc123"
	fmt.Println(StringEndsWith(testString, "123"))
}

func TestSubstring(t *testing.T) {
	testString := "abc123"
	endIndex := strings.Index(testString, "123")
	fmt.Println(Substring(testString, 0, endIndex))
}

func TestDelete(t *testing.T) {
	var slice = []int{0, 1, 2, 3}
	slice = slices.Delete(slice, 2, 3)
	fmt.Println(slice)
}
