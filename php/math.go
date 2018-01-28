package php

import (
	"math"
	"math/cmplx"
	"math/rand"
	"strconv"
	"time"
)

//Absolute xue
func Abs(x float64) float64 {

	return math.Abs(x)
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

//Decimal to binary
func Decbin(x int64) string {

	return strconv.FormatInt(x, 2)
}

//Decimal to hexadecimal
func Dechex(x int64) string {

	return strconv.FormatInt(x, 16)
}

//Decimal to octal
func Decoct(x int64) string {

	return strconv.FormatInt(x, 8)
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

//Finds whether a value is a legal finite number
func IsFinite(f float64, sign int) bool {

	return !math.IsInf(f, sign)
}

//Finds whether a value is infinite
func IsInfinite(f float64, sign int) bool {

	return math.IsInf(f, sign)
}

//Finds whether a value is not a number
func IsNan(f float64) bool {

	return math.IsNaN(f)
}

//Natural logarithm
func Log(x float64) float64 {

	return math.Log(x)
}

//Base-10 logarithm
func Log10(x float64) float64 {

	return math.Log10(x)
}

// Returns log(1 + number)
// computed in a way that is accurate even when the value of number is close to zero
func Log1p(x float64) float64 {

	return math.Log1p(x)
}

//Find highest value
func Max(x, y float64) float64 {

	return math.Max(x, y)
}

//Find lowest value
func Min(x, y float64) float64 {

	return math.Min(x, y)
}

//Get value of pi
func Pi() float64 {

	return math.Pi
}

//Exponential expression
func Pow(x, y float64) float64 {

	return math.Pow(x, y)
}

//Generate a random integer
func Rand(args ...int) int {

	rand.Seed(time.Now().Unix())
	l := len(args)
	if l > 1 {
		min := math.Min(float64(args[0]), float64(args[1]))
		max := math.Max(float64(args[0]), float64(args[1]))
		return int(min + rand.Float64()*(max-min))
	} else if l > 0 {
		return rand.Intn(args[0])
	} else {
		return rand.Intn(1 << 32)
	}
}

//Rounds a float
func Round(x float64) float64 {

	return math.Floor(x + 0.5)
}

//Sine
func Sin(x float64) float64 {

	return math.Sin(x)
}

//Hyperbolic sine
func Sinh(x float64) float64 {

	return math.Sinh(x)
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
