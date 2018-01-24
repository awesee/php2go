package php

import "os"

//Change directory
func Chdir(dir string) error {

	return os.Chdir(dir)
}
