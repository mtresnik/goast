package functions

import (
	"goast/pkg/operations"
	"math/cmplx"
	"strconv"
)

type Division struct {
	Numerator   operations.Operation
	Denominator operations.Operation
}

func (d Division) GetValues() []operations.Operation {
	return []operations.Operation{d.Numerator, d.Denominator}
}

func (d Division) IsConstant() bool {
	if d.Numerator.IsConstant() && d.Denominator.IsConstant() {
		return true
	}
	if operations.Equals(d.Numerator, d.Denominator) {
		return true
	}
	return false
}

func (d Division) ToString() string {
	if d.IsConstant() {
		c := d.ToNumber()
		return strconv.FormatComplex(c, 'f', 5, 64)
	}
	retString := ""
	retString += d.Numerator.ToString()
	retString += "/"
	retString += d.Denominator.ToString()
	return retString
}

func (d Division) ToNumber() complex128 {
	if d.Numerator.IsConstant() && d.Denominator.IsConstant() {
		var numerator = d.Numerator.ToNumber()
		var denominator = d.Denominator.ToNumber()
		if numerator == 0 && denominator == 0 {
			return cmplx.NaN()
		}
		if denominator == 0 {
			return cmplx.Inf()
		}
		return numerator / denominator
	}
	if operations.Equals(d.Numerator, d.Denominator) {
		return 1
	}
	return cmplx.NaN()
}

func (d Division) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Division{
		Numerator:   d.Numerator.Evaluate(one, other),
		Denominator: d.Denominator.Evaluate(one, other),
	}
}
