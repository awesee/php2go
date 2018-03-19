package php

// ArrayFlip - Exchanges all keys with their associated values in an array
func ArrayFlip(arrayMap map[interface{}]interface{}) map[interface{}]interface{} {

	r := make(map[interface{}]interface{})
	for i, v := range arrayMap {
		r[v] = i
	}

	return r
}
