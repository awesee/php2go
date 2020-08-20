package php

import "os"

// Chdir - Change directory
func Chdir(dir string) error {

	return os.Chdir(dir)
}

// Getcwd - Get current directory
func Getcwd() (dir string) {

	dir, err := os.Getwd()
	if err != nil {
		dir = err.Error()
	}
	return
}
