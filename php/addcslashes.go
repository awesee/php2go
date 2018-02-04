package php

// Addcslashes - Quote string with slashes in a C style
func Addcslashes(s string, c byte) string {

	l := len(s)
	r := make([]byte, l, l*2)
	for i := 0; i < l; i++ {
		if s[i] == c {
			r = append(r, '\\')
		}

		r = append(r, s[i])
	}

	return string(r)
}
