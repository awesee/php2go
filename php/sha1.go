package php

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

//Sha1 - Calculate the sha1 hash of a string
func Sha1(s string) string {

	byte := sha1.Sum([]byte(s))

	return hex.EncodeToString(byte[:])
}

//Sha1File - Calculates the md5 hash of a given file
func Sha1File(filename string) string {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	byte := sha1.Sum(data)

	return hex.EncodeToString(byte[:])
}
