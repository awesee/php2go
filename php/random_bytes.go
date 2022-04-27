package php

import "crypto/rand"

func RandomBytes(len int) []byte {
	b := make([]byte, len)
	_, _ = rand.Read(b)
	return b
}
