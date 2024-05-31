package parser

import (
	"fmt"
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/parser/tokens"
	"goast/pkg/operations/variables"
	"goast/pkg/utils"
	"slices"
	"sort"
	"strings"
)

func Parse(inputString string) (operations.Operation, *error) {
	err := validateString(inputString)
	if err != nil {
		return constants.NaN, err
	}
	tokenList := tokens.Tokenize(inputString)
	err = validateSyntax(tokenList)
	if err != nil {
		return constants.NaN, err
	}
	var intermediateOperation Intermediate
	intermediateOperation, err = generateIntermediate(tokenList)
	if err != nil {
		return constants.NaN, err
	}
	operation := intermediateOperation.Compile()
	operation = operation.Evaluate(variables.E, constants.E)
	operation = operation.Evaluate(variables.I, constants.I)
	return operation, nil
}

func validateString(inputString string) *error {
	balance := 0
	for _, v := range inputString {
		if v == '(' {
			balance--
		} else if v == ')' {
			balance++
		}
	}
	if balance != 0 {
		err := fmt.Errorf("imbalanced number of parentheses")
		return &err
	}
	var accumulate = ""
	for _, v := range inputString {
		if v == tokens.Decimal && strings.ContainsRune(accumulate, tokens.Decimal) {
			err := fmt.Errorf("too many decimals for given input string")
			return &err
		}
		if strings.ContainsRune(tokens.ValidNumbers, v) == false {
			accumulate = ""
		} else {
			accumulate += string(v)
		}
	}
	return nil
}

func validateSyntax(tokenList []tokens.Token) *error {
	for i := 0; i < len(tokenList)-1; i++ {
		var curr = tokenList[i]
		if curr.TokenType == tokens.Operator {
			if (*curr.Representation)[0] == '+' || (*curr.Representation)[0] == '-' {
				var next = tokenList[i+1]
				if next.TokenType == tokens.ClosedParentheses || next.TokenType == tokens.Operator {
					err := fmt.Errorf("invalid syntax at: %s\t%s", curr.ToString(), next.ToString())
					return &err
				}
			}
		}
	}
	return nil
}

const (
	number      = tokens.Number
	parentheses = tokens.OpenParentheses
	function    = tokens.Function
	variable    = tokens.Variable
)

type tokenSet struct {
	StartIndex     int
	EndIndex       int
	TokenSetType   int
	Tokens         []tokens.Token
	Representation *string
}

func sortByStartIndex(tokenList []tokenSet) {
	sort.Slice(tokenList, func(i, j int) bool {
		return tokenList[i].StartIndex < tokenList[j].StartIndex
	})
}

func equals(one tokenSet, other tokenSet) bool {
	if one.StartIndex != other.StartIndex {
		return false
	}
	if one.EndIndex != other.EndIndex {
		return false
	}
	if one.TokenSetType != other.TokenSetType {
		return false
	}
	return tokens.ContainsAll(one.Tokens, other.Tokens)
}

func generateMultipleIntermediates(tokenList []tokens.Token) ([]Intermediate, *error) {
	var tokenSets, err = generateTokenSets(tokenList)
	if err != nil {
		return make([]Intermediate, 0), err
	}
	var intermediates []Intermediate
	intermediates, err = generateIntermediates(tokenSets)
	intermediates = generateOperators(intermediates, tokenList)
	return intermediates, nil
}

func generateIntermediate(tokenList []tokens.Token) (Intermediate, *error) {
	intermediates, err := generateMultipleIntermediates(tokenList)
	if err != nil {
		return NullIntermediate, err
	}
	if len(intermediates) == 1 {
		return intermediates[0], nil
	}
	err1 := fmt.Errorf("error generating intermediates")
	return NullIntermediate, &err1
}

func generateTokenSets(inputList []tokens.Token) ([]tokenSet, *error) {
	var err *error
	var tokenSets = generateParentheses(inputList)
	tokenSets, err = generateFunctions(tokenSets, inputList)
	if err != nil {
		return make([]tokenSet, 0), err
	}
	tokenSets = generateVariables(tokenSets, inputList)
	return generateNumbers(tokenSets, inputList), nil
}

func generateParentheses(inputList []tokens.Token) []tokenSet {
	var retList = make([]tokenSet, 0)
	var inner = make([]tokens.Token, 0)
	var balance = 0
	var startIndex = -1
	for i, token := range inputList {
		if token.TokenType == tokens.OpenParentheses {
			balance--
		} else if token.TokenType == tokens.ClosedParentheses {
			balance++
		}
		if balance == -1 && token.TokenType == tokens.OpenParentheses {
			startIndex = i
		}
		if balance == 0 && token.TokenType == tokens.ClosedParentheses {
			for j := startIndex + 1; j < i; j++ {
				inner = append(inner, inputList[j])
			}
			retList = append(retList, tokenSet{
				StartIndex:   startIndex,
				EndIndex:     i,
				TokenSetType: parentheses,
				Tokens:       slices.Clone(inner),
			})
			startIndex = -1
			inner = make([]tokens.Token, 0)
		}
	}
	return retList
}

func generateFunctions(current []tokenSet, inputList []tokens.Token) ([]tokenSet, *error) {
	var clone = slices.Clone(current)
	var retList = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedToken(i, current) == false {
			if token.TokenType == tokens.Function {
				var found *tokenSet = nil
				var foundIndex = -1
				var expectedIndex = i + 1
				for j, set := range clone {
					if set.TokenSetType == parentheses {
						if set.StartIndex == expectedIndex {
							found = &set
							foundIndex = j
							break
						}
					}
				}
				if found == nil {
					err := fmt.Errorf("could not find parentheses for given function: %s", *(token.Representation))
					return make([]tokenSet, 0), &err
				}
				clone = slices.Delete(clone, foundIndex, foundIndex+1)
				for j := 0; j < len(retList); j++ {
					if equals(*found, retList[j]) {
						foundIndex = j
						break
					}
				}
				retList = slices.Delete(retList, foundIndex, foundIndex+1)
				var fn = tokenSet{
					StartIndex:     i,
					EndIndex:       found.EndIndex,
					TokenSetType:   function,
					Tokens:         (*found).Tokens,
					Representation: token.Representation,
				}
				retList = append(retList, fn)
			}
		}
	}
	return retList, nil
}

func generateVariables(current []tokenSet, inputList []tokens.Token) []tokenSet {
	var retList = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedToken(i, current) == false {
			if token.TokenType == tokens.Variable {
				var variable = tokenSet{
					StartIndex:     i,
					EndIndex:       i,
					TokenSetType:   variable,
					Tokens:         []tokens.Token{token},
					Representation: token.Representation,
				}
				retList = append(retList, variable)
			}
		}
	}
	sortByStartIndex(retList)
	return retList
}

func generateNumbers(current []tokenSet, inputList []tokens.Token) []tokenSet {
	var retList = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedToken(i, current) == false {
			if token.TokenType == tokens.Number {
				var number = tokenSet{
					StartIndex:   i,
					EndIndex:     i,
					TokenSetType: number,
					Tokens:       []tokens.Token{token},
				}
				retList = append(retList, number)
			}
		}
	}
	sortByStartIndex(retList)
	return retList
}

func generateIntermediates(current []tokenSet) ([]Intermediate, *error) {
	var retList = make([]Intermediate, 0)
	for _, set := range current {
		switch set.TokenSetType {
		case number:
			retList = append(retList, INumber{
				StartIndex: set.StartIndex,
				EndIndex:   set.EndIndex,
				Token:      set.Tokens[0],
			})
			break
		case variable:
			retList = append(retList, IVariable{
				StartIndex: set.StartIndex,
				EndIndex:   set.EndIndex,
				Token:      set.Tokens[0],
			})
			break
		case parentheses:
			inner, err := generateIntermediate(set.Tokens)
			if err != nil {
				return make([]Intermediate, 0), err
			}
			retList = append(retList, IParentheses{
				StartIndex: set.StartIndex,
				EndIndex:   set.EndIndex,
				Inner:      inner,
			})
			break
		case function:
			inner, err := generateMultipleIntermediates(set.Tokens)
			if err != nil {
				return make([]Intermediate, 0), err
			}
			retList = append(retList, IFunction{
				StartIndex: set.StartIndex,
				EndIndex:   set.EndIndex,
				Name:       *(set.Representation),
				Inner:      inner,
			})
			break
		default:
			break
		}
	}
	return retList, nil
}

func generateOperators(current []Intermediate, inputList []tokens.Token) []Intermediate {
	intermediates := generateIdentities(current, inputList)
	intermediates = generatePowers(intermediates, inputList)
	intermediates = generateMultiplicationAndDivision(intermediates, inputList)
	intermediates = generateAdditionAndSubtraction(intermediates, inputList)
	return intermediates
}

func generateIdentities(current []Intermediate, inputList []tokens.Token) []Intermediate {
	var clone = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedOperation(i, clone) == false {
			if token.TokenType == tokens.Operator {
				if (*token.Representation) == tokens.Plus {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					// TODO add error handling for when right is null
					if left == nil && right != nil {
						clone = Remove(clone, *right)
						clone = append(clone, IIdentity{
							StartIndex: i,
							EndIndex:   (*right).GetEndIndex(),
							Inner:      *right,
						})
					}
				} else if (*token.Representation) == tokens.Minus {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					// TODO add error handling for when right is null
					if left == nil && right != nil {
						clone = Remove(clone, *right)
						clone = append(clone, INegation{
							StartIndex: i,
							EndIndex:   (*right).GetEndIndex(),
							Inner:      *right,
						})
					}
				}
			}
		}
	}
	var retList = slices.Clone(clone)
	SortByStartIndex(retList)
	return retList
}

func generatePowers(current []Intermediate, inputList []tokens.Token) []Intermediate {
	var clone = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedOperation(i, clone) == false {
			if token.TokenType == tokens.Operator {
				if (*token.Representation) == "^" {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					// TODO Require left and right for power
					if left != nil && right != nil {
						clone = Remove(clone, *left)
						clone = Remove(clone, *right)
						clone = append(clone, IPower{
							StartIndex: (*left).GetStartIndex(),
							EndIndex:   (*right).GetEndIndex(),
							Left:       *left,
							Right:      *right,
						})
					}
				}
			}
		}
	}
	var retList = slices.Clone(clone)
	SortByStartIndex(retList)
	return retList
}

func generateMultiplicationAndDivision(current []Intermediate, inputList []tokens.Token) []Intermediate {
	// TODO Require left and right checks for mult and div
	var clone = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedOperation(i, clone) == false {
			if token.TokenType == tokens.Operator {
				if (*token.Representation) == "*" {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					if left != nil && right != nil {
						clone = Remove(clone, *left)
						clone = Remove(clone, *right)
						clone = append(clone, IMultiplication{
							StartIndex: (*left).GetStartIndex(),
							EndIndex:   (*right).GetEndIndex(),
							Left:       *left,
							Right:      *right,
						})
					}
				} else if (*token.Representation) == "/" {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					if left != nil && right != nil {
						clone = Remove(clone, *left)
						clone = Remove(clone, *right)
						clone = append(clone, IDivision{
							StartIndex: (*left).GetStartIndex(),
							EndIndex:   (*right).GetEndIndex(),
							Left:       *left,
							Right:      *right,
						})
					}
				}
			}
		}
	}
	var retList = slices.Clone(clone)
	SortByStartIndex(retList)
	return retList
}

func generateAdditionAndSubtraction(current []Intermediate, inputList []tokens.Token) []Intermediate {
	// TODO Require left and right checks for add and sub
	var clone = slices.Clone(current)
	for i, token := range inputList {
		if indexProcessedOperation(i, clone) == false {
			if token.TokenType == tokens.Operator {
				if (*token.Representation) == tokens.Plus {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					if left != nil && right != nil {
						clone = Remove(clone, *left)
						clone = Remove(clone, *right)
						clone = append(clone, IAddition{
							StartIndex: (*left).GetStartIndex(),
							EndIndex:   (*right).GetEndIndex(),
							Left:       *left,
							Right:      *right,
						})
					}
				} else if (*token.Representation) == tokens.Minus {
					var left = getLeftIntermediate(i, clone)
					var right = getRightIntermediate(i, clone)
					if left != nil && right != nil {
						clone = Remove(clone, *left)
						clone = Remove(clone, *right)
						clone = append(clone, ISubtraction{
							StartIndex: (*left).GetStartIndex(),
							EndIndex:   (*right).GetEndIndex(),
							Left:       *left,
							Right:      *right,
						})
					}
				}
			}
		}
	}

	var retList = slices.Clone(clone)
	SortByStartIndex(retList)
	return retList
}

func indexProcessedToken(i int, tokenList []tokenSet) bool {
	for _, t := range tokenList {
		if utils.IntInRangeInclusive(i, t.StartIndex, t.EndIndex) {
			return true
		}
	}
	return false
}

func indexProcessedOperation(i int, intermediateList []Intermediate) bool {
	for _, t := range intermediateList {
		if utils.IntInRangeInclusive(i, t.GetStartIndex(), t.GetEndIndex()) {
			return true
		}
	}
	return false
}

func getLeftIntermediate(i int, intermediateList []Intermediate) *Intermediate {
	for _, intermediate := range intermediateList {
		if utils.IntInRangeInclusive(i-1, intermediate.GetStartIndex(), intermediate.GetEndIndex()) {
			return &intermediate
		}
	}
	return nil
}

func getRightIntermediate(i int, intermediateList []Intermediate) *Intermediate {
	for _, intermediate := range intermediateList {
		if utils.IntInRangeInclusive(i+1, intermediate.GetStartIndex(), intermediate.GetEndIndex()) {
			return &intermediate
		}
	}
	return nil
}
