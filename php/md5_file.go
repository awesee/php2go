package php

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

//Calculates the md5 hash of a given file
func Md5File(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	byte := md5.Sum(data)

	return hex.EncodeToString(byte[:]), nil
}
