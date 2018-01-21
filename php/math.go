package php

import (
	"math"
	"math/cmplx"
)

//Absolute xue
func Abs(x float64) {
	math.Abs(x)
}

//Arc cosine
func Acos(x complex128) complex128 {
	return cmplx.Acos(x)
}

//Inverse hyperbolic cosine
func Acosh(x complex128) complex128 {
	return cmplx.Acosh(x)
}

//Arc sine
func Asin(x complex128) complex128 {
	return cmplx.Asin(x)
}

//Inverse hyperbolic sine
func Asinh(x complex128) complex128 {
	return cmplx.Asinh(x)
}

//Arc tangent
func Atan(x complex128) complex128 {
	return cmplx.Atan(x)
}

//Inverse hyperbolic tangent
func Atanh(x complex128) complex128 {
	return cmplx.Atanh(x)
}
