package php

import "os"

// SysGetTempDir - Returns directory path used for temporary files
func SysGetTempDir() string {

	return os.TempDir()
}
