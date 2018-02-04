package php

import (
	"os"
	"syscall"
)

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

// Closedir - Close directory's handle
func Closedir(fd int) (err error) {

	return syscall.Close(fd)
}
