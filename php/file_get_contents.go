package php

import "io/ioutil"

// FileGetContents - Reads entire file into a string
func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
