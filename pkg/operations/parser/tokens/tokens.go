package tokens

import (
	"fmt"
	"github.com/mtresnik/goast/pkg/utils"
	"sort"
)

const (
	Number = iota
	Operator
	OpenParenthesis
	ClosedParenthesis
	Text
	Function
	Variable
)

const (
	Plus      = "+"
	PlusRune  = '+'
	Minus     = "-"
	MinusRune = '-'
)

type Token struct {
	StartIndex     int
	EndIndex       int
	TokenType      int
	Representation *string
}

func (t Token) Convert(other int) Token {
	return Token{
		StartIndex:     t.StartIndex,
		EndIndex:       t.EndIndex,
		TokenType:      other,
		Representation: t.Representation,
	}
}

func (t Token) ToString() string {
	Representation := "nil"
	if t.Representation != nil {
		Representation = *t.Representation
	}
	return fmt.Sprintf("[range:[%d,%d],TokenType:%d,Representation:%s]", t.StartIndex, t.EndIndex, t.TokenType, Representation)
}

func SingleIndex(index int, TokenType int) Token {
	return Token{
		StartIndex:     index,
		EndIndex:       index,
		TokenType:      TokenType,
		Representation: nil,
	}
}

func NullIndex(TokenType int, rep *string) Token {
	return Token{
		StartIndex:     -1,
		EndIndex:       -1,
		TokenType:      TokenType,
		Representation: rep,
	}
}

func ToString(tokenList []Token) string {
	retString := "["
	for i, v := range tokenList {
		retString += v.ToString()
		if i < len(tokenList)-1 {
			retString += ", "
		}
	}
	retString += "]"
	return retString
}

func SortByStartIndex(tokenList []Token) {
	sort.Slice(tokenList, func(i, j int) bool {
		return tokenList[i].StartIndex < tokenList[j].StartIndex
	})
}

func IndexProcessed(index int, tokenList []Token) bool {
	for _, t := range tokenList {
		if utils.IntInRangeInclusive(index, t.StartIndex, t.EndIndex) {
			return true
		}
	}
	return false
}

func ContainsToken(tokenList []Token, other Token) bool {
	for _, token := range tokenList {
		if Equals(token, other) {
			return true
		}
	}
	return false
}

func ContainsAll(oneList []Token, otherList []Token) bool {
	if len(oneList) != len(otherList) {
		return false
	}
	for _, token := range oneList {
		if ContainsToken(otherList, token) == false {
			return false
		}
	}
	return true
}

func Equals(one Token, other Token) bool {
	if one.StartIndex != other.StartIndex {
		return false
	}
	if one.EndIndex != other.EndIndex {
		return false
	}
	if one.Representation == nil && other.Representation != nil {
		return false
	}
	if other.Representation == nil && one.Representation != nil {
		return false
	}
	if one.Representation == nil && other.Representation == nil {
		return true
	}
	oneRep, otherRep := *(one.Representation), *(other.Representation)
	return oneRep == otherRep
}
