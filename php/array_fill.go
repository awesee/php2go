package php

// ArrayFill - Fill an array with values
func ArrayFill(startIndex uint, num uint, value interface{}) map[uint]interface{} {

	r := make(map[uint]interface{})
	var i uint
	for i = 0; i < num; i++ {
		r[startIndex] = value
		startIndex++
	}

	return r
}
