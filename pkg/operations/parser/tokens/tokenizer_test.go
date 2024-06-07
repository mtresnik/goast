package tokens

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	var inputString = "(()"
	var _ = Tokenize(inputString)
}

func TestTokenize2(t *testing.T) {
	var inputString = "123.5 * 3xyz + (456)"
	var result = Tokenize(inputString)
	fmt.Println(ToString(result))
}

func TestTokenize3(t *testing.T) {
	var inputString = "123.5 * 3xyzsin(5) + (456) + sin"
	var result = Tokenize(inputString)
	fmt.Println(ToString(result))
}

func TestTokenize4(t *testing.T) {
	var inputString = "123.5 * 3xyz * sin(5) + + + + (456) + sin"
	var result = Tokenize(inputString)
	fmt.Println(ToString(result))
}

func TestTokenize5(t *testing.T) {
	var inputString = "log_(x,y)"
	var result = Tokenize(inputString)
	fmt.Println(ToString(result))
}
