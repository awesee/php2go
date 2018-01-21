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

//Arc tangent of two variables
func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
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

//Hyperbolic cosine
func Cosh(x float64) float64 {
	return math.Cosh(x)
}

//Calculates the exponent of e
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Returns exp(number) - 1
// computed in a way that is accurate even when the value of number is close to zero
func Expm1(x float64) float64 {
	return math.Exp(x) - 1
}

//Round fractions down
func Floor(x float64) float64 {
	return math.Floor(x)
}

//Natural logarithm
func Log(x float64) float64 {
	return math.Log(x)
}

//Find highest value
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

//Find lowest value
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

//Square root
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

//Tangent
func Tan(x float64) float64 {
	return math.Tan(x)
}

//Hyperbolic tangent
func Tanh(x float64) float64 {
	return math.Tanh(x)
}
