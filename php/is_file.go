package php

import "os"

// IsFile Tells whether the filename is a regular file
func IsFile(name string) bool {

	var exist = true
	if _, err := os.Stat(name); os.IsNotExist(err) {
		exist = false
	}

	return exist
}
