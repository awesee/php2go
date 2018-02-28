package php

import "encoding/base64"

// Base64Encode - Encodes data with MIME base64
func Base64Encode(s string) string {

	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode - Decodes data encoded with MIME base64
func Base64Decode(s string) (string, error) {

	bt, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}
