package php

// ArrayCountValues — Counts all the values of an array
func ArrayCountValues(s []interface{}) map[interface{}]uint {

	r := make(map[interface{}]uint)
	for _, v := range s {
		if c, ok := r[v]; ok {
			r[v] = c + 1
		} else {
			r[v] = 1
		}
	}

	return r
}
