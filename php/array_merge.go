package php

// ArrayMerge — Merge one or more arrays
func ArrayMerge(arr ...[]interface{}) []interface{} {

	s := make([]interface{}, 0)
	for _, v := range arr {
		s = append(s, v...)
	}

	return s
}
