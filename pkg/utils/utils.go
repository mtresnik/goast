package utils

import (
	"reflect"
	"strings"
	"unicode/utf8"
)

func unpackArray(s any) []any {
	v := reflect.ValueOf(s)
	r := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Index(i).Interface()
	}
	return r
}

func hasArrays(root []any) bool {
	for i := 0; i < len(root); i++ {
		if reflect.TypeOf(root[i]) == reflect.TypeOf(root) {
			return true
		}
	}
	return false
}

func flatten(root []any) []any {
	var count = 0
	for i := 0; i < len(root); i++ {
		if reflect.TypeOf(root[i]) == reflect.TypeOf(root) {
			childArray := unpackArray(root[i])
			count += len(childArray)
		} else {
			count++
		}
	}
	var retArray = make([]any, count)
	var outerIndex = 0
	for i := 0; i < len(root); i++ {
		if reflect.TypeOf(root[i]) == reflect.TypeOf(root) {
			childArray := unpackArray(root[i])
			for j := 0; j < len(childArray); j++ {
				retArray[outerIndex] = childArray[j]
				outerIndex++
			}
		} else {
			retArray[outerIndex] = root[i]
			outerIndex++
		}
	}
	return retArray
}

func deepFlatten(root []any) []any {
	var retArray = root
	for hasArrays(retArray) {
		retArray = flatten(retArray)
	}
	return retArray
}

func MinMaxInt(one int, two int) (minValue int, maxValue int) {
	if one < two {
		minValue = one
		maxValue = two
	} else {
		minValue = two
		maxValue = one
	}
	return
}

func IntInRangeInclusive(test int, one int, two int) bool {
	minValue, maxValue := MinMaxInt(one, two)
	if test >= minValue && test <= maxValue {
		return true
	}
	return false
}

func StringStartsWith(one string, other string) bool {
	return strings.Index(one, other) == 0
}

func StringEndsWith(one string, other string) bool {
	lastIndex := utf8.RuneCountInString(one)
	return strings.LastIndex(one, other) == lastIndex-utf8.RuneCountInString(other)
}

func SubstringToEnd(input string, startIndex int) string {
	return Substring(input, startIndex, utf8.RuneCountInString(input))
}

func Substring(input string, startIndex int, endIndex int) string {
	runeSlice := []rune(input)
	if startIndex >= len(runeSlice) {
		return ""
	}
	var length = endIndex - startIndex
	if endIndex > len(runeSlice) {
		length = len(runeSlice) - startIndex
		endIndex = startIndex + length
	}
	return string(runeSlice[startIndex:endIndex])
}

func FindRemainingStrings(test string, key string) []string {
	if utf8.RuneCountInString(key) == 0 {
		return make([]string, 0)
	}
	if strings.Contains(test, key) == false {
		return make([]string, 0)
	}
	if utf8.RuneCountInString(test) == utf8.RuneCountInString(key) {
		return make([]string, 0)
	}
	var index = strings.Index(test, key)
	if index == 0 {
		return []string{SubstringToEnd(test, utf8.RuneCountInString(key))}
	}
	if StringEndsWith(test, key) {
		return []string{Substring(test, 0, utf8.RuneCountInString(test)-utf8.RuneCountInString(key))}
	} else {
		return []string{
			Substring(test, 0, index),
			SubstringToEnd(test, index+utf8.RuneCountInString(key)),
		}
	}
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
