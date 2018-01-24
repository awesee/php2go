package php

import "os"

//Returns directory path used for temporary files
func SysGetTempDir() string {

	return os.TempDir()
}
