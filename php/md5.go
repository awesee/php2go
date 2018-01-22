package php

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

//Calculate the md5 hash of a string
func Md5(s string) string {
	byte := md5.Sum([]byte(s))
	return hex.EncodeToString(byte[:])
}

//Calculates the md5 hash of a given file
func Md5File(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	byte := md5.Sum(data)

	return hex.EncodeToString(byte[:]), nil
}
