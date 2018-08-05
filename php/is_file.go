package php

import "os"

// IsFile Tells whether the filename is a regular file
func IsFile(name string) bool {

	fi, err := os.Stat(name)

	return err == nil && fi.Mode().IsRegular()
}

