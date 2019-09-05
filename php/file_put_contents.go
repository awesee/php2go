package php

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// FilePutContents - Write data to a file
func FilePutContents(filename string, data []byte) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}
