package goast

import (
	"github.com/mtresnik/goutils/pkg/goutils"
	"math/cmplx"
)

type Power struct {
	Base     Operation
	Exponent Operation
}

func (p Power) GetValues() []Operation {
	return []Operation{p.Base, p.Exponent}
}

func (p Power) IsConstant() bool {
	if p.Exponent.IsConstant() && p.Exponent.Number() == 0+0i {
		return true
	}
	return p.Base.IsConstant() && p.Exponent.IsConstant()
}

func (p Power) String() string {
	if p.IsConstant() {
		c := p.Number()
		return goutils.SmartComplexString(c)
	}
	retString := p.Base.String()
	retString += " ^ "
	retString += p.Exponent.String()
	return retString
}

func (p Power) Number() complex128 {
	if p.IsConstant() == false {
		return cmplx.NaN()
	}
	if p.Exponent.Number() == 0+0i {
		if p.Base.IsConstant() && p.Base.Number() == 0+0i {
			return cmplx.NaN()
		}
		return 1
	}
	if p.Base.Number() == 0+0i {
		if p.Exponent.IsConstant() {
			exponentNumber := p.Exponent.Number()
			if exponentNumber == 0+0i {
				return cmplx.NaN()
			}
			if imag(exponentNumber) == 0 && real(exponentNumber) < 0 {
				return cmplx.Inf()
			}
		}
		return 0
	}
	return cmplx.Pow(p.Base.Number(), p.Exponent.Number())
}

func (p Power) Evaluate(one Operation, other Operation) Operation {
	if OperationEquals(p, one) {
		return other
	}
	retArray := EvaluateValues(p, one, other)
	return Power{Base: retArray[0], Exponent: retArray[1]}
}
