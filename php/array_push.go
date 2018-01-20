package php

//Push one or more elements onto the end of array
func ArrayPush(s *[]string, args ...string) {
	for _, v := range args {
		*s = append(*s, v)
	}
}
