package php

import (
	"encoding/json"
)

//Returns the JSON representation of a value
func JsonEncode(data []byte, v interface{}) error {

	return json.Unmarshal(data, v)
}

//Decodes a JSON string
func JsonDecode(v interface{}) ([]byte, error) {

	return json.Marshal(v)
}
