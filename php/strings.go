package php

import (
	"strconv"
	"strings"
)

//Convert binary data into hexadecimal representation
func Bin2hex(b string) string {

	base, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return ""
	}

	return strconv.FormatInt(base, 16)
}

//Binary to decimal
func Bindec(b string) int64 {

	i, _ := strconv.ParseInt(b, 2, 64)

	return i
}

//Decodes a hexadecimally encoded binary string
func Hex2bin(x string) string {

	base, err := strconv.ParseInt(x, 16, 64)
	if err != nil {
		return ""
	}

	return strconv.FormatInt(base, 2)
}

//Return a specific character
func Chr(ascii int) string {

	for ascii < 0 {
		ascii += 256
	}
	ascii %= 256

	return string(ascii)
}

//Return ASCII value of character
func Ord(s byte) byte {

	return s
}

//Strip whitespace (or other characters) from the beginning and end of a string
func Trim(s, cutset string) string {

	return strings.Trim(s, cutset)
}

//Find the position of the first occurrence of a substring in a string
func Strpos(s, substr string) int {

	return strings.Index(s, substr)
}

//Find the position of the first occurrence of a case-insensitive substring in a string
func Stripos(s, substr string) int {

	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.Index(s, substr)
}

// Find the position of the last occurrence of a substring in a string
func Strrpos(s, substr string) int {

	return strings.LastIndex(s, substr)
}

//Find the position of the last occurrence of a case-insensitive substring in a string
func Strripos(s, substr string) int {

	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.LastIndex(s, substr)
}

//Find the last occurrence of a character in a string
func Strrchr(s, substr string) string {

	i := strings.LastIndex(s, substr)
	if i < 0 {
		return ""
	}

	return s[i:]
}
