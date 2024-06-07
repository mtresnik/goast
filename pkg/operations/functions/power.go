package functions

import (
	"goast/pkg/operations"
	"goast/pkg/utils"
	"math/cmplx"
)

type Power struct {
	Base     operations.Operation
	Exponent operations.Operation
}

func (p Power) GetValues() []operations.Operation {
	return []operations.Operation{p.Base, p.Exponent}
}

func (p Power) IsConstant() bool {
	if p.Exponent.IsConstant() && p.Exponent.ToNumber() == 0+0i {
		return true
	}
	return p.Base.IsConstant() && p.Exponent.IsConstant()
}

func (p Power) ToString() string {
	if p.IsConstant() {
		c := p.ToNumber()
		return utils.SmartComplexString(c)
	}
	retString := p.Base.ToString()
	retString += " ^ "
	retString += p.Exponent.ToString()
	return retString
}

func (p Power) ToNumber() complex128 {
	if p.IsConstant() == false {
		return cmplx.NaN()
	}
	if p.Exponent.ToNumber() == 0+0i {
		if p.Base.IsConstant() && p.Base.ToNumber() == 0+0i {
			return cmplx.NaN()
		}
		return 1
	}
	if p.Base.ToNumber() == 0+0i {
		if p.Exponent.IsConstant() {
			exponentNumber := p.Exponent.ToNumber()
			if exponentNumber == 0+0i {
				return cmplx.NaN()
			}
			if imag(exponentNumber) == 0 && real(exponentNumber) < 0 {
				return cmplx.Inf()
			}
		}
		return 0
	}
	return cmplx.Pow(p.Base.ToNumber(), p.Exponent.ToNumber())
}

func (p Power) Evaluate(one operations.Operation, other operations.Operation) operations.Operation {
	if operations.Equals(p, one) {
		return other
	}
	retArray := operations.EvaluateValues(p, one, other)
	return Power{Base: retArray[0], Exponent: retArray[1]}
}
