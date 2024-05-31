package tokens

import (
	"fmt"
	"testing"
)

const (
	First = iota
	Second
	Third
)

func TestToken_ToString(t *testing.T) {
	var Representation = "a"
	var token1 = Token{
		StartIndex:     0,
		EndIndex:       1,
		TokenType:      First,
		Representation: &Representation,
	}
	fmt.Println(token1.ToString())
}

func TestToken_ToString2(t *testing.T) {
	var token1 = Token{
		StartIndex:     0,
		EndIndex:       1,
		TokenType:      Third,
		Representation: nil,
	}
	fmt.Println(token1.ToString())
}
