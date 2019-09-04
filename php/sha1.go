package php

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

// Sha1 - Calculate the sha1 hash of a string
func Sha1(s string) string {
	digest := sha1.Sum([]byte(s))
	return hex.EncodeToString(digest[:])
}

// Sha1File - Calculates the md5 hash of a given file
func Sha1File(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	digest := sha1.Sum(data)
	return hex.EncodeToString(digest[:])
}
