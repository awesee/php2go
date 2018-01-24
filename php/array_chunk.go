package php

import (
	"fmt"
	"math"
)

type ArraySlice []interface{}

func ArrayChunk(input ArraySlice, size int) ArraySlice {

	length := len(input)
	count := int(math.Ceil(float64(length) / float64(size)))
	ret := make(ArraySlice, count)
	for i := 0; i < count-1; i++ {
		ret[i] = input[i*size : (i+1)*size]
		fmt.Println(ret)
	}
	ret[count-1] = input[size*(count-1):]

	return ret
}
