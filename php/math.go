package php

import (
	"math"
	"math/cmplx"
	"strconv"
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

//Convert a number between arbitrary bases
func BaseConvert(num string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(num, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

//Round fractions up
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

//Cosine
func Cos(x float64) float64 {
	return math.Cos(x)
}
