package parser

import (
	"fmt"
	"goast/pkg/operations"
	"goast/pkg/operations/constants"
	"goast/pkg/operations/functions"
	"goast/pkg/operations/parser/tokens"
	"goast/pkg/operations/variables"
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
