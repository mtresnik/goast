package functions

import (
	"goast/pkg/operations"
	"goast/pkg/utils"
	"math/cmplx"
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
		return utils.SmartComplexString(c)
	}
	retString := ""
	retString += d.Numerator.ToString()
	retString += "/"
	retString += d.Denominator.ToString()
	return retString
}

func (d Division) ToNumber() complex128 {
	if operations.Equals(d.Numerator, d.Denominator) {
		return 1
	}
	if d.IsConstant() == false {
		return cmplx.NaN()
	}
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
	return cmplx.NaN()
}

func (d Division) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	return Division{
		Numerator:   d.Numerator.Evaluate(one, other),
		Denominator: d.Denominator.Evaluate(one, other),
	}
}
