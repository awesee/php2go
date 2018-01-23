package php

import (
	"strings"
)

//Find the position of the first occurrence of a substring in a string
func Strpos(s, substr string) int {

	return strings.Index(s, substr)
}

//Find the last occurrence of a character in a string
func Strrchr(s, substr string) string {

	i := strings.LastIndex(s, substr)
	if i < 0 {
		return ""
	}

	return s[i:]
}
