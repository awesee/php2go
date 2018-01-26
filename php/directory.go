package php

import (
	"os"
	"syscall"
)

//Change directory
func Chdir(dir string) error {

	return os.Chdir(dir)
}

// +------------------------------------------------------------
// | @desc      Get current directory
// | @param
// | @return    string, error
// |
// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------

func Getcwd() (dir string) {

	dir, err := os.Getwd()
	if err != nil {
		dir = err.Error()
	}
	return
}

// +------------------------------------------------------------
// | @desc      Close directory's handle
// | @param
// | @return    error
// |
// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Closedir(fd int) (err error) {

	return syscall.Close(fd)
}
