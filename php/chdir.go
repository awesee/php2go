package php

import "os"

func Chdir(dir string) error {
	return os.Chdir(dir)
}
