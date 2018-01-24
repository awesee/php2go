package php

import (
	"strings"
)

//Returns trailing name component of path
func Basename(s string, suffix ...string) string {

	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if len(suffix) > 0 {
		suf := strings.LastIndex(s, suffix[0])
		s = s[:suf]
	}

	return s
}
