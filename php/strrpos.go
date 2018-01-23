package php

import "strings"

// Find the position of the last occurrence of a substring in a string
func Strrpos(s, substr string) int {

	return strings.LastIndex(s, substr)
}
