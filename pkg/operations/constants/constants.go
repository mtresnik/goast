package constants

import (
	"math"
	"math/cmplx"
)

var eString = "e"
var iString = "i"

var E = Constant{Representation: math.E, StringRepresentation: &eString}
var PI = Constant{Representation: math.Pi}
var I = Constant{Representation: 1i, StringRepresentation: &iString}
var NaN = Constant{Representation: cmplx.NaN()}
var ZERO = Constant{Representation: 0}
var ONE = Constant{Representation: 1}
var TWO = Constant{Representation: 2}
var TEN = Constant{Representation: 10}

func New(representation complex128) Constant {
	return Constant{Representation: representation}
}
