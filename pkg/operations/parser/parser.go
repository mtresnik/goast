package parser

import (
	"fmt"
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/functions"
	"goast/pkg/operations/parser/tokens"
	"goast/pkg/operations/variables"
	"goast/pkg/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Intermediate interface {
	GetStartIndex() int
	GetEndIndex() int
	GetTokens() []tokens.Token
	Compile() operations.Operation
	ToString() string
}

func Equals(one Intermediate, other Intermediate) bool {
	if one.GetStartIndex() != other.GetStartIndex() {
		return false
	}
	if one.GetEndIndex() != other.GetEndIndex() {
		return false
	}
	return tokens.ContainsAll(one.GetTokens(), other.GetTokens())
}

func getTokens(intermediates []Intermediate) []tokens.Token {
	retList := make([]tokens.Token, 0)
	for _, op := range intermediates {
		retList = slices.Concat(retList, op.GetTokens())
	}
	return retList
}

func compile(intermediates []Intermediate) []operations.Operation {
	retList := make([]operations.Operation, 0)
	for _, op := range intermediates {
		retList = append(retList, op.Compile())
	}
	return retList
}

// INumber <editor-fold desc="INumber">
type INumber struct {
	StartIndex int
	EndIndex   int
	Token      tokens.Token
}

func (n INumber) GetStartIndex() int {
	return n.StartIndex
}

func (n INumber) GetEndIndex() int {
	return n.EndIndex
}

func (n INumber) GetTokens() []tokens.Token {
	return []tokens.Token{n.Token}
}

func (n INumber) Compile() operations.Operation {
	if n.Token.Representation == nil {
		return constants.NaN
	}
	representation, err := strconv.ParseComplex(strings.Replace(*(n.Token.Representation), " ", "", -1), 64)
	if err != nil {
		return constants.NaN
	}
	return constants.Constant{Representation: representation}
}

func (n INumber) ToString() string {
	return fmt.Sprintf("[INumber, range:[%d, %d], representation:%s]", n.StartIndex, n.EndIndex, *(n.Token.Representation))
}

// </editor-fold>

// IVariable <editor-fold desc="IVariable">
type IVariable struct {
	StartIndex int
	EndIndex   int
	Token      tokens.Token
}

func (v IVariable) GetStartIndex() int {
	return v.StartIndex
}

func (v IVariable) GetEndIndex() int {
	return v.EndIndex
}

func (v IVariable) GetTokens() []tokens.Token {
	return []tokens.Token{v.Token}
}

func (v IVariable) Compile() operations.Operation {
	return variables.Variable{Name: *(v.Token.Representation)}
}

func (v IVariable) ToString() string {
	return fmt.Sprintf("[IVariable, range:[%d, %d], representation:%s]", v.StartIndex, v.EndIndex, *(v.Token.Representation))
}

// </editor-fold>

// IAddition <editor-fold desc="IAddition">
type IAddition struct {
	StartIndex int
	EndIndex   int
	Left       Intermediate
	Right      Intermediate
}

func (a IAddition) GetStartIndex() int {
	return a.StartIndex
}

func (a IAddition) GetEndIndex() int {
	return a.EndIndex
}

func (a IAddition) GetTokens() []tokens.Token {
	return getTokens([]Intermediate{a.Left, a.Right})
}

func (a IAddition) Compile() operations.Operation {
	return functions.Addition{
		Values: []operations.Operation{a.Left.Compile(), a.Right.Compile()},
	}
}

func (a IAddition) ToString() string {
	return fmt.Sprintf("[IAddition, range:[%d, %d], representation:%s]", a.StartIndex, a.EndIndex, tokens.ToString(a.GetTokens()))
}

// </editor-fold>

// ISubtraction <editor-fold desc="ISubtraction">
type ISubtraction struct {
	StartIndex int
	EndIndex   int
	Left       Intermediate
	Right      Intermediate
}

func (s ISubtraction) GetStartIndex() int {
	return s.StartIndex
}

func (s ISubtraction) GetEndIndex() int {
	return s.EndIndex
}

func (s ISubtraction) GetTokens() []tokens.Token {
	return getTokens([]Intermediate{s.Left, s.Right})
}

func (s ISubtraction) Compile() operations.Operation {
	return functions.Subtraction{
		Values: []operations.Operation{s.Left.Compile(), s.Right.Compile()},
	}
}

func (s ISubtraction) ToString() string {
	return fmt.Sprintf("[ISubtraction, range:[%d, %d], representation:%s]", s.StartIndex, s.EndIndex, tokens.ToString(s.GetTokens()))
}

// </editor-fold>

// IDivision <editor-fold desc="IDivision">
type IDivision struct {
	StartIndex int
	EndIndex   int
	Left       Intermediate
	Right      Intermediate
}

func (d IDivision) GetStartIndex() int {
	return d.StartIndex
}

func (d IDivision) GetEndIndex() int {
	return d.EndIndex
}

func (d IDivision) GetTokens() []tokens.Token {
	return getTokens([]Intermediate{d.Left, d.Right})
}

func (d IDivision) Compile() operations.Operation {
	return functions.Division{
		Numerator:   d.Left.Compile(),
		Denominator: d.Right.Compile(),
	}
}

func (d IDivision) ToString() string {
	return fmt.Sprintf("[IDivision, range:[%d, %d], representation:%s]", d.StartIndex, d.EndIndex, tokens.ToString(d.GetTokens()))
}

// </editor-fold>

// IIdentity <editor-fold desc="IIdentity">
type IIdentity struct {
	StartIndex int
	EndIndex   int
	Inner      Intermediate
}

func (i IIdentity) GetStartIndex() int {
	return i.StartIndex
}

func (i IIdentity) GetEndIndex() int {
	return i.EndIndex
}

func (i IIdentity) GetTokens() []tokens.Token {
	return i.Inner.GetTokens()
}

func (i IIdentity) Compile() operations.Operation {
	return i.Inner.Compile()
}

func (i IIdentity) ToString() string {
	return fmt.Sprintf("[IIdentity, range:[%d, %d], representation:%s]", i.StartIndex, i.EndIndex, i.Inner.ToString())
}

// </editor-fold>

// INegation <editor-fold desc="INegation">
type INegation struct {
	StartIndex int
	EndIndex   int
	Inner      Intermediate
}

func (n INegation) GetStartIndex() int {
	return n.StartIndex
}

func (n INegation) GetEndIndex() int {
	return n.EndIndex
}

func (n INegation) GetTokens() []tokens.Token {
	return n.Inner.GetTokens()
}

func (n INegation) Compile() operations.Operation {
	return functions.Negation{Inner: n.Inner.Compile()}
}

func (n INegation) ToString() string {
	return fmt.Sprintf("[INegation, range:[%d, %d], representation:%s]", n.StartIndex, n.EndIndex, n.Inner.ToString())
}

// </editor-fold>

// IMultiplication <editor-fold desc="IMultiplication">
type IMultiplication struct {
	StartIndex int
	EndIndex   int
	Left       Intermediate
	Right      Intermediate
}

func (m IMultiplication) GetStartIndex() int {
	return m.StartIndex
}

func (m IMultiplication) GetEndIndex() int {
	return m.EndIndex
}

func (m IMultiplication) GetTokens() []tokens.Token {
	return getTokens([]Intermediate{m.Left, m.Right})
}

func (m IMultiplication) Compile() operations.Operation {
	return functions.Multiplication{
		Values: []operations.Operation{m.Left.Compile(), m.Right.Compile()},
	}
}

func (m IMultiplication) ToString() string {
	return fmt.Sprintf("[IMultiplication, range:[%d, %d], representation:%s]", m.StartIndex, m.EndIndex, tokens.ToString(m.GetTokens()))
}

// </editor-fold>

// IParentheses <editor-fold desc="IParentheses">
type IParentheses struct {
	StartIndex int
	EndIndex   int
	Inner      Intermediate
}

func (p IParentheses) GetStartIndex() int {
	return p.StartIndex
}

func (p IParentheses) GetEndIndex() int {
	return p.EndIndex
}

func (p IParentheses) GetTokens() []tokens.Token {
	return p.Inner.GetTokens()
}

func (p IParentheses) Compile() operations.Operation {
	return functions.Parentheses{Inner: p.Inner.Compile()}
}

func (p IParentheses) ToString() string {
	return fmt.Sprintf("[IParentheses, range:[%d, %d], representation:%s]", p.StartIndex, p.EndIndex, p.Inner.ToString())
}

// </editor-fold>

// IPower <editor-fold desc="IPower">
type IPower struct {
	StartIndex int
	EndIndex   int
	Left       Intermediate
	Right      Intermediate
}

func (p IPower) GetStartIndex() int {
	return p.StartIndex
}

func (p IPower) GetEndIndex() int {
	return p.EndIndex
}

func (p IPower) GetTokens() []tokens.Token {
	return getTokens([]Intermediate{p.Left, p.Right})
}

func (p IPower) Compile() operations.Operation {
	return functions.Power{Base: p.Left.Compile(), Exponent: p.Right.Compile()}
}

func (p IPower) ToString() string {
	return fmt.Sprintf("[IPower, range:[%d, %d], representation:%s]", p.StartIndex, p.EndIndex, tokens.ToString(p.GetTokens()))
}

// </editor-fold>

// IFunction <editor-fold desc="IFunction">
type IFunction struct {
	StartIndex int
	EndIndex   int
	Name       string
	Inner      []Intermediate
}

func (f IFunction) GetStartIndex() int {
	return f.StartIndex
}

func (f IFunction) GetEndIndex() int {
	return f.EndIndex
}

func (f IFunction) GetTokens() []tokens.Token {
	return getTokens(f.Inner)
}

func (f IFunction) Compile() operations.Operation {
	return functions.BuildFunction(f.Name, compile(f.Inner)...)
}

func (f IFunction) ToString() string {
	return fmt.Sprintf("[IFunction, range:[%d, %d], representation:%s]", f.StartIndex, f.EndIndex, tokens.ToString(f.GetTokens()))
}

// </editor-fold>

var NullIntermediate = INumber{
	StartIndex: -1,
	EndIndex:   -1,
	Token: tokens.Token{
		StartIndex:     -1,
		EndIndex:       -1,
		TokenType:      tokens.Number,
		Representation: nil,
	},
}

func SortByStartIndex(intermediates []Intermediate) {
	sort.Slice(intermediates, func(i, j int) bool {
		return intermediates[i].GetStartIndex() < intermediates[j].GetStartIndex()
	})
}

func IndexOf(intermediates []Intermediate, other Intermediate) int {
	for i, value := range intermediates {
		if Equals(value, other) {
			return i
		}
	}
	return -1
}

func Remove(intermediates []Intermediate, other Intermediate) []Intermediate {
	index := IndexOf(intermediates, other)
	if index != -1 {
		return slices.Delete(intermediates, index, index+1)
	}
	return intermediates
}

func ParseOperation(inputString string) (operations.Operation, *error) {
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
				if next.TokenType == tokens.ClosedParenthesis || next.TokenType == tokens.Operator {
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
	parentheses = tokens.OpenParenthesis
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
	if err != nil {
		return make([]Intermediate, 0), err
	}
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
		if token.TokenType == tokens.OpenParenthesis {
			balance--
		} else if token.TokenType == tokens.ClosedParenthesis {
			balance++
		}
		if balance == -1 && token.TokenType == tokens.OpenParenthesis {
			startIndex = i
		}
		if balance == 0 && token.TokenType == tokens.ClosedParenthesis {
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
	var clone = current
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
	var retList = clone
	SortByStartIndex(retList)
	return retList
}

func generatePowers(current []Intermediate, inputList []tokens.Token) []Intermediate {
	var clone = current
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
	var retList = clone
	SortByStartIndex(retList)
	return retList
}

func generateMultiplicationAndDivision(current []Intermediate, inputList []tokens.Token) []Intermediate {
	// TODO Require left and right checks for mult and div
	var clone = current
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
	var retList = clone
	SortByStartIndex(retList)
	return retList
}

func generateAdditionAndSubtraction(current []Intermediate, inputList []tokens.Token) []Intermediate {
	// TODO Require left and right checks for add and sub
	var clone = current
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

	var retList = clone
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
