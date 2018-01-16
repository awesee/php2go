package php

import "os"

func IsFile(name string) bool {
	var exist = true
	if _, err := os.Stat(name); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
