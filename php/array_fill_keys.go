package php

// ArrayFillKeys - Fill an array with values, specifying keys
func ArrayFillKeys(keys []interface{}, value interface{}) map[interface{}]interface{} {

	r := make(map[interface{}]interface{})
	for _, v := range keys {
		r[v] = value
	}

	return r
}
