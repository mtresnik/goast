package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Division struct {
	Numerator   Operation
	Denominator Operation
}

func (d Division) GetValues() []Operation {
	return []Operation{d.Numerator, d.Denominator}
}

func (d Division) IsConstant() bool {
	if d.Numerator.IsConstant() && d.Denominator.IsConstant() {
		return true
	}
	if OperationEquals(d.Numerator, d.Denominator) {
		return true
	}
	return false
}

func (d Division) String() string {
	if d.IsConstant() {
		c := d.Number()
		return goutils.SmartComplexString(c)
	}
	retString := ""
	retString += d.Numerator.String()
	retString += "/"
	retString += d.Denominator.String()
	return retString
}

func (d Division) Number() complex128 {
	if OperationEquals(d.Numerator, d.Denominator) {
		return 1
	}
	if d.IsConstant() == false {
		return cmplx.NaN()
	}
	if d.Numerator.IsConstant() && d.Denominator.IsConstant() {
		var numerator = d.Numerator.Number()
		var denominator = d.Denominator.Number()
		if numerator == 0 && denominator == 0 {
			return cmplx.NaN()
		}
		if denominator == 0 {
			return cmplx.Inf()
		}
		return numerator / denominator
	}
	return cmplx.NaN()
}

func (d Division) Evaluate(one Operation, other Operation) Operation {
	return Division{
		Numerator:   d.Numerator.Evaluate(one, other),
		Denominator: d.Denominator.Evaluate(one, other),
	}
}
