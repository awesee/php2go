package php

import "encoding/base64"

//Encodes data with MIME base64
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
