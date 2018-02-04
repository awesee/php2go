package php

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

//Md5 - Calculate the md5 hash of a string
func Md5(s string) string {

	byte := md5.Sum([]byte(s))

	return hex.EncodeToString(byte[:])
}

//Md5File - Calculates the md5 hash of a given file
func Md5File(filename string) string {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	byte := md5.Sum(data)

	return hex.EncodeToString(byte[:])
}
