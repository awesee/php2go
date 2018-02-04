package php

// ArrayReverse - Return an array with elements in reverse order
func ArrayReverse(s []interface{}) []interface{} {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
