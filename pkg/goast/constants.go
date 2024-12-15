package goast

import (
	"math"
	"math/cmplx"
)

var eString = "e"
var iString = "i"

var E_Constant = Constant{Representation: math.E, StringRepresentation: &eString}
var PI_Constant = Constant{Representation: math.Pi}
var I_Constant = Constant{Representation: 1i, StringRepresentation: &iString}
var NaN_Constant = Constant{Representation: cmplx.NaN()}
var ZERO_Constant = Constant{Representation: 0}
var ONE_Constant = Constant{Representation: 1}
var TWO_Constant = Constant{Representation: 2}
var TEN_Constant = Constant{Representation: 10}

func NewConstant(representation complex128) Constant {
	return Constant{Representation: representation}
}
