package php

import (
	"encoding/json"
)

// JsonEncode - Returns the JSON representation of a value
func JsonEncode(v interface{}) ([]byte, error) {

	return json.Marshal(v)
}

// JsonDecode - Decodes a JSON string
func JsonDecode(data []byte, v interface{}) error {

	return json.Unmarshal(data, v)
}
