package php

import (
	"html"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Constants for StrPad
const (
	StrPadRight = "STR_PAD_RIGHT"
	StrPadLeft  = "STR_PAD_LEFT"
)

// Bin2hex - Convert binary data into hexadecimal representation
func Bin2hex(b string) string {
	base, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(base, 16)
}

// Bindec - Binary to decimal
func Bindec(b string) int64 {
	i, _ := strconv.ParseInt(b, 2, 64)
	return i
}

// Hex2bin - Decodes a hexadecimally encoded binary string
func Hex2bin(x string) string {
	base, err := strconv.ParseInt(x, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(base, 2)
}

// Chr - Return a specific character
func Chr(ascii int) string {
	for ascii < 0 {
		ascii += 256
	}
	return string(ascii % 256)
}

// Ord - Return ASCII value of character
func Ord(s byte) byte {
	return s
}

// Explode - Split a string by string
func Explode(s, sep string) []string {
	return strings.Split(s, sep)
}

// GetHtmlTranslationTable - Returns the translation table used by htmlspecialchars() and htmlentities()
func GetHtmlTranslationTable() map[string]string {
	return map[string]string{
		`"`: "&quot;",
		`&`: "&amp;",
		`<`: "&lt;",
		`>`: "&gt;",
	}
}

// Htmlspecialchars - Convert special characters to HTML entities
func Htmlspecialchars(s string) string {
	return html.EscapeString(s)
}

// HtmlspecialcharsDecode - Convert special HTML entities back to characters
func HtmlspecialcharsDecode(s string) string {
	return html.UnescapeString(s)
}

// Implode - Join array elements with a string
func Implode(a []string, sep string) string {
	return strings.Join(a, sep)
}

// Join - Alias of implode()
func Join(a []string, sep string) string {
	return Implode(a, sep)
}

//StripTags - Strip HTML and PHP tags from a string
func StripTags(s string) string {
	reg, _ := regexp.Compile(`<[\S\s]+?>`)
	s = reg.ReplaceAllStringFunc(s, strings.ToLower)
	//remove style
	reg, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	s = reg.ReplaceAllString(s, "")
	//remove script
	reg, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	s = reg.ReplaceAllString(s, "")

	reg, _ = regexp.Compile(`<[\S\s]+?>`)
	s = reg.ReplaceAllString(s, "\n")

	reg, _ = regexp.Compile(`\s{2,}`)
	s = reg.ReplaceAllString(s, "\n")

	return strings.TrimSpace(s)
}

// Trim - Strip whitespace (or other characters) from the beginning and end of a string
func Trim(s, cutset string) string {
	if cutset == "" {
		return strings.TrimSpace(s)
	}
	return strings.Trim(s, cutset)
}

// Ltrim - Strip whitespace (or other characters) from the beginning of a string
func Ltrim(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// Rtrim - Strip whitespace (or other characters) from the end of a string
func Rtrim(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// Nl2br - Inserts HTML line breaks before all newlines in a string
func Nl2br(s string) string {
	return strings.ReplaceAll(s, "\n", "\n<br />")
}

// StrPad - Pad a string to a certain length with another string
func StrPad(s string, length int, args ...string) string {
	runes := []rune(s)
	l := len(runes)
	if l > length {
		return s
	}
	padString := " "
	padType := StrPadRight
	if len(args) > 1 {
		padString = args[0]
		padType = args[1]
	} else if len(args) > 0 {
		padString = args[0]
	}

	padStringLen := len([]rune(padString))
	count := (length-l)/padStringLen + 1
	out := ""
	padString = strings.Repeat(padString, count)
	if padType == StrPadLeft {
		out = string([]rune(padString)[:length-l]) + s
	} else {
		out = s + string([]rune(padString)[:length-l])
	}
	return out
}

// StrRepeat - Repeat a string
func StrRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// StrReplace - Replace all occurrences of the search string with the replacement string
func StrReplace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Strtolower - Make a string lowercase
func Strtolower(s string) string {
	return strings.ToLower(s)
}

// Strtoupper - Make a string uppercase
func Strtoupper(s string) string {
	return strings.ToUpper(s)
}

// Strstr - Find the first occurrence of a string
func Strstr(s, substr string) int {
	return strings.Index(s, substr)
}

// Strpos - Find the position of the first occurrence of a substring in a string
func Strpos(s, substr string) int {
	return strings.Index(s, substr)
}

// Stripos - Find the position of the first occurrence of a case-insensitive substring in a string
func Stripos(s, substr string) int {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)

	return strings.Index(s, substr)
}

// Strrpos - Find the position of the last occurrence of a substring in a string
func Strrpos(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// Strripos - Find the position of the last occurrence of a case-insensitive substring in a string
func Strripos(s, substr string) int {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.LastIndex(s, substr)
}

// Strrchr - Find the last occurrence of a character in a string
func Strrchr(s, substr string) string {
	i := strings.LastIndex(s, substr)
	if i < 0 {
		return ""
	}
	return s[i:]
}

// Strlen - Get string length
func Strlen(s string) int {
	return len(s)
}

// MbStrlen - Get string length
func MbStrlen(s string) int {
	return utf8.RuneCountInString(s)
}

// Strrev - Reverse a string
func Strrev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Substr - Return part of a string
func Substr(s string, start int, length ...int) string {
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(s) + l
			if end < 0 {
				end = 0
			}
			return s[start:end]
		}
		end := start + l
		if end > len(s) {
			end = len(s)
		}
		return s[start:end]
	}
	return s[start:]
}

// MbSubstr - Get part of string
func MbSubstr(s string, start int, length ...int) string {
	runes := []rune(s)
	if len(length) > 0 {
		l := length[0]
		if l < 0 {
			end := len(runes) + l
			if end < 0 {
				end = 0
			}
			return string(runes[start:end])
		}
		end := start + l
		if end > len(runes) {
			end = len(runes)
		}
		return string(runes[start:end])
	}
	return string(runes[start:])
}

// SubstrCount - Count the number of substring occurrences
func SubstrCount(s, substr string) int {
	return strings.Count(s, substr)
}

// Ucfirst - Make a string's first character uppercase
func Ucfirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Ucwords — Uppercase the first character of each word in a string
func Ucwords(s string) string {
	return strings.Title(s)
}
