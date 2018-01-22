package php

import (
	"crypto/sha1"
	"encoding/hex"
)

//Calculate the sha1 hash of a string
func Sha1(s string) string {
	byte := sha1.Sum([]byte(s))
	return hex.EncodeToString(byte[:])
}
