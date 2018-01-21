package php

import "os"

//Changes file mode
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}
