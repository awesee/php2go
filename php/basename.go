package php

import "path/filepath"

// Basename - Returns trailing name component of path
func Basename(path string) string {
	return filepath.Base(path)
}
