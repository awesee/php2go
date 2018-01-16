package php

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	byte := md5.Sum([]byte(s))
	return hex.EncodeToString(byte[:])
}
