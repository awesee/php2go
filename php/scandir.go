package php

import (
	"io/ioutil"
	"os"
)

// scandir — List files and directories inside the specified path
func Scandir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}
