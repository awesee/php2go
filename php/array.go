package php

//Create an array
func Array(v ...interface{}) []interface{} {

	return v
}

//Count all elements in an array, or something in an object
func Count(v []interface{}) int {

	return len(v)
}
