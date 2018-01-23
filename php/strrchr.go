package php

import (
	"strings"
)

//Find the last occurrence of a character in a string
func Strrchr(s, substr string) string {

	i := strings.LastIndex(s, substr)
	if i < 0 {
		return ""
	}

	return s[i:]
}
