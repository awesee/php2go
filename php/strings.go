package php

import (
	"html"
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

//Split a string by string
func Explode(s, sep string) []string {

	if s == "" {
		return []string{s}
	}

	return strings.Split(s, sep)
}

//Returns the translation table used by htmlspecialchars() and htmlentities()
func GetHtmlTranslationTable() map[string]string {

	return map[string]string{
		`"`: "&quot;",
		`&`: "&amp;",
		`<`: "&lt;",
		`>`: "&gt;",
	}
}

//Convert special characters to HTML entities
func Htmlspecialchars(s string) string {

	return html.EscapeString(s)
}

//Convert special HTML entities back to characters
func HtmlspecialcharsDecode(s string) string {

	return html.UnescapeString(s)
}

//Join array elements with a string
func Implode(a []string, sep string) string {

	return strings.Join(a, sep)
}

//Alias of implode()
func Join(a []string, sep string) string {

	return Implode(a, sep)
}

//Strip whitespace (or other characters) from the beginning and end of a string
func Trim(s, cutset string) string {

	return strings.Trim(s, cutset)
}

//Strip whitespace (or other characters) from the beginning of a string
func Ltrim(s, cutset string) string {

	return strings.TrimLeft(s, cutset)
}

//Strip whitespace (or other characters) from the end of a string
func Rtrim(s, cutset string) string {

	return strings.TrimRight(s, cutset)
}

//Make a string lowercase
func Strtolower(s string) string {

	return strings.ToLower(s)
}

//Make a string uppercase
func Strtoupper(s string) string {

	return strings.ToUpper(s)
}

//Find the first occurrence of a string
func Strstr(s, substr string) int {

	return strings.Index(s, substr)
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

//Get string length
func Strlen(s string) int {

	return len(s)
}

//Get string length
func MbStrlen(s string) int {

	return len([]rune(s))
}

//Reverse a string
func Strrev(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

//Return part of a string
func Substr(s string, start int, length ...int) string {

	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(s) + l
			return string(s[start:end])
		} else {
			end := start + l
			return string(s[start:end])
		}
	}

	return s[start:]
}

//Get part of string
func MbSubstr(s string, start int, length ...int) string {

	runes := []rune(s)
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(runes) + l
			return string(runes[start:end])
		} else {
			end := start + l
			return string(runes[start:end])
		}
	}

	return string(runes[start:])
}

//Make a string's first character uppercase
func Ucfirst(s string) string {

	runes := []rune(s)
	if len(runes) < 1 {
		return s
	}

	return strings.ToUpper(string(runes[0])) + string(runes[1:])
}
