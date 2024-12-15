package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"slices"
	"strings"
	"unicode/utf8"
)

const (
	Decimal                    = '.'
	ValidNumbers               = "0123456789."
	Operators                  = "+-*/^,"
	OpenParenthesisCharacter   = '('
	ClosedParenthesisCharacter = ')'
)

func Tokenize(inputStringParam string) []Token {
	inputString := preProcess(inputStringParam)
	tokenList := tokenizeNumbers(inputString)
	tokenList = tokenizeOperators(tokenList, inputString)
	tokenList = tokenizeParentheses(tokenList, inputString)
	tokenList = tokenizeText(tokenList, inputString)
	tokenList = tokenizeFunctions(tokenList)
	tokenList = tokenizeVariables(tokenList)
	return postProcess(tokenList)
}

func preProcess(inputString string) string {
	retString := strings.ReplaceAll(inputString, " ", "")
	return retString
}

func tokenizeNumbers(inputString string) []Token {
	retArray := make([]Token, utf8.RuneCountInString(inputString))
	count := 0
	var accumulate = ""
	for i, v := range inputString {
		if strings.ContainsRune(ValidNumbers, v) == false {
			if utf8.RuneCountInString(accumulate) != 0 {
				var start = i - utf8.RuneCountInString(accumulate)
				var end = i - 1
				var representation = strings.Clone(accumulate)
				var number = Token{
					StartIndex:     start,
					EndIndex:       end,
					TokenType:      NumberToken,
					Representation: &representation,
				}
				retArray[count] = number
				count++
				accumulate = ""
			}
		} else {
			accumulate += string(v)
		}
	}
	if utf8.RuneCountInString(accumulate) != 0 {
		var start = utf8.RuneCountInString(inputString) - utf8.RuneCountInString(accumulate)
		var end = utf8.RuneCountInString(inputString) - 1
		var representation = strings.Clone(accumulate)
		var number = Token{
			StartIndex:     start,
			EndIndex:       end,
			TokenType:      NumberToken,
			Representation: &representation,
		}
		retArray[count] = number
		count++
	}
	return retArray[:count]
}

func tokenizeOperators(tokenList []Token, inputString string) []Token {
	retList := tokenList
	for i, v := range inputString {
		if IndexProcessed(i, retList) == false {
			if strings.ContainsRune(Operators, v) {
				var representation = strings.Clone(string(v))
				var operator = SingleIndex(i, OperatorToken)
				operator.Representation = &representation
				retList = append(retList, operator)
			}
		}
	}
	return retList
}

func tokenizeParentheses(tokenList []Token, inputString string) []Token {
	retList := tokenList
	for i, v := range inputString {
		if IndexProcessed(i, retList) == false {
			if v == OpenParenthesisCharacter {
				var representation = strings.Clone(string(v))
				var parenthesis = SingleIndex(i, OpenParenthesisToken)
				parenthesis.Representation = &representation
				retList = append(retList, parenthesis)
			} else if v == ClosedParenthesisCharacter {
				var representation = strings.Clone(string(v))
				var parenthesis = SingleIndex(i, ClosedParenthesisToken)
				parenthesis.Representation = &representation
				retList = append(retList, parenthesis)
			}
		}
	}
	return retList
}

func tokenizeText(tokenList []Token, inputString string) []Token {
	retList := tokenList
	var accumulated = ""
	for i, v := range inputString {
		if IndexProcessed(i, retList) {
			if utf8.RuneCountInString(accumulated) > 0 {
				var start = i - utf8.RuneCountInString(accumulated)
				var end = i - 1
				var representation = strings.Clone(accumulated)
				var text = Token{
					StartIndex:     start,
					EndIndex:       end,
					TokenType:      TextToken,
					Representation: &representation,
				}
				retList = append(retList, text)
			}
			accumulated = ""
		} else {
			accumulated += string(v)
		}
	}
	if utf8.RuneCountInString(accumulated) > 0 {
		var start = utf8.RuneCountInString(inputString) - utf8.RuneCountInString(accumulated)
		var end = utf8.RuneCountInString(inputString) - 1
		var representation = strings.Clone(accumulated)
		var text = Token{
			StartIndex:     start,
			EndIndex:       end,
			TokenType:      TextToken,
			Representation: &representation,
		}
		retList = append(retList, text)
	}
	TokenSortByStartIndex(retList)
	return retList
}

func tokenizeFunctions(tokenList []Token) []Token {
	var retList = make([]Token, 0)
	for i, curr := range tokenList {
		if curr.TokenType != TextToken {
			retList = append(retList, curr)
		} else {
			if i < len(tokenList)-1 && (tokenList[i+1]).TokenType == OpenParenthesisToken {
				var innerFunc = ""
				var foundInner = false
				var representation = *curr.Representation
				for _, key := range ReservedFunctions {
					if utf8.RuneCountInString(key) <= utf8.RuneCountInString(representation) &&
						goutils.StringEndsWith(representation, key) {
						innerFunc = key
						foundInner = true
						break
					}
				}
				if foundInner {
					var endIndex = strings.LastIndex(representation, innerFunc)
					if endIndex != 0 {
						var newRep = goutils.Substring(representation, 0, endIndex)
						var rem = NullIndex(TextToken, &newRep)
						retList = append(retList, rem)
					}
					var function = NullIndex(FunctionToken, &innerFunc)
					retList = append(retList, function)
				} else {
					var rem = curr.Convert(TextToken)
					retList = append(retList, rem)
				}
			} else {
				var rem = curr.Convert(TextToken)
				retList = append(retList, rem)
			}
		}
	}
	return retList
}

func tokenizeVariables(tokenList []Token) []Token {
	var retList = make([]Token, 0)
	for _, curr := range tokenList {
		if curr.TokenType != TextToken {
			retList = append(retList, curr)
		} else {
			retList = slices.Concat(retList, maxVariablesInString(*curr.Representation))
		}
	}
	return retList
}

func maxVariablesInString(inputString string) []Token {
	var retList = make([]Token, 0)
	var maxVar = ""
	var maxCount = -1
	for _, key := range ReservedFunctions {
		if strings.Contains(inputString, key) && utf8.RuneCountInString(key) > maxCount {
			maxVar = key
			maxCount = utf8.RuneCountInString(maxVar)
		}
	}
	var remainingStrings = goutils.FindRemainingStrings(inputString, maxVar)
	if maxCount == -1 || len(remainingStrings) == 0 {
		retList = append(retList, NullIndex(VariableToken, &inputString))
		return retList
	}
	if len(remainingStrings) == 1 {
		// Right case
		if goutils.StringStartsWith(inputString, maxVar) {
			retList = append(retList, NullIndex(VariableToken, &maxVar))
			retList = slices.Concat(retList, maxVariablesInString(remainingStrings[0]))
			return retList
		}
		// Left case
		var leftHandSide = maxVariablesInString(remainingStrings[0])
		retList = slices.Concat(retList, leftHandSide)
		retList = append(retList, NullIndex(VariableToken, &maxVar))
		return retList
	}
	// Middle case
	if len(remainingStrings) == 2 {
		var leftHandSide = maxVariablesInString(remainingStrings[0])
		var rightHandSide = maxVariablesInString(remainingStrings[1])
		retList = slices.Concat(retList, leftHandSide)
		retList = append(retList, NullIndex(VariableToken, &maxVar))
		retList = slices.Concat(retList, rightHandSide)
		return retList
	}
	retList = append(retList, NullIndex(VariableToken, &inputString))
	return retList
}

func postProcess(tokenList []Token) []Token {
	var currList = justifyMultiplication(tokenList)
	var collapsed = collapseSigns(currList)
	for len(collapsed) != len(currList) {
		currList = collapsed
		collapsed = collapseSigns(currList)
	}
	return currList
}

// justifyMultiplication converts instances of [5,x] to [5,*,x]
func justifyMultiplication(inputList []Token) []Token {
	var retList = make([]Token, 0)
	for i, curr := range inputList {
		retList = append(retList, curr)
		if (curr.TokenType == NumberToken || curr.TokenType == VariableToken || curr.TokenType == ClosedParenthesisToken) &&
			i < len(inputList)-1 &&
			inputList[i+1].TokenType != OperatorToken &&
			inputList[i+1].TokenType != ClosedParenthesisToken {
			var representation = "*"
			retList = append(retList, NullIndex(OperatorToken, &representation))
		}
	}
	return retList
}

func collapseSigns(inputList []Token) []Token {
	var retList = make([]Token, 0)
	var i = 0
	for i < len(inputList) {
		var curr = inputList[i]
		var currRepresentation = *(curr.Representation)
		var nextRepresentation = ""
		if i < len(inputList)-1 {
			nextRepresentation = *(inputList[i+1].Representation)
		}
		if curr.TokenType == OperatorToken &&
			i < len(inputList)-1 &&
			inputList[i+1].TokenType == OperatorToken &&
			(currRepresentation[0] == '+' || currRepresentation[0] == '-') &&
			(nextRepresentation[0] == '+' || nextRepresentation[0] == '-') {
			var representation = Plus
			if currRepresentation[0] != nextRepresentation[0] {
				representation = Minus
			}
			retList = append(retList, NullIndex(OperatorToken, &representation))
			i += 2
		} else {
			retList = append(retList, curr)
			i++
		}
	}

	return retList
}
