package php

// Addslashes - Quote string with slashes
func Addslashes(s string) string {

	l := len(s)
	r := make([]byte, l, l*2)
	for i := 0; i < l; i++ {

		switch s[i] {
		case '\'', '"', '\\':
			r = append(r, '\\')
		}

		r = append(r, s[i])
	}

	return string(r)
}
