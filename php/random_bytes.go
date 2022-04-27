package php

import "crypto/rand"

// RandomBytes — Generates cryptographically secure pseudo-random bytes
func RandomBytes(len int) []byte {
	b := make([]byte, len)
	_, _ = rand.Read(b)
	return b
}
