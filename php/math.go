package php

import (
	"math"
	"math/cmplx"
	"strconv"
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

// +------------------------------------------------------------
// | @desc      Decimal To Bytes
// | @param     n int
// | @return    []byte

// | @since     https://studygolang.com/articles/1599

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

// +------------------------------------------------------------
// | @desc      Byte To Hex
// | @param     input []byte, hex_count int
// | @return    []byte

// | @since     https://studygolang.com/articles/1599

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func intToHex(input []byte, hex_count int) []byte {
	out := make([]byte, len(input)+int(hex_count))
	copy(out, input)
	for i := 0; i < hex_count; i++ {
		out[len(input)+i] = byte(hex_count)
	}
	return out
}

//Decimal to binary

// +------------------------------------------------------------
// | @desc      Decimal to binary
// | @param     x int64
// | @return    []byte

// | @since     https://studygolang.com/articles/1599

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Decbin(x int64) []byte {

	return intToHex(IntToBytes(x), 2)
}

//Decimal to hexadecimal
// +------------------------------------------------------------
// | @desc      Decimal to hexadecimal
// | @param     x int64
// | @return    []byte

// | @since     https://studygolang.com/articles/1599

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Dechex(x int64) string {

	return intToHex(IntToBytes(x), 16)
}

//Decimal to octal
// +------------------------------------------------------------
// | @desc      Decimal to octal
// | @param     x int64
// | @return    []byte

// | @since     https://studygolang.com/articles/1599

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Decoct(x int64) string {

	return intToHex(IntToBytes(x), 8)
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
