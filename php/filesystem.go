package php

import "os"

//Changes file mode
func Chmod(name string, mode os.FileMode) error {

	return os.Chmod(name, mode)
}

// +------------------------------------------------------------
// | @desc      make directort
// | @param
// | @return    error
// |
// | @author    Openset <jinheking@gmail.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Mkdir(name string, mode os.FileMode) error {

	return os.Mkdir(name, mode)
}

// +------------------------------------------------------------
// | @desc      is directort
// | @param
// | @return    bool , error
// |
// | @author    Openset <jinheking@gmail.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func IsDir(name string) (b bool, err error) {

	fd, err := os.Stat(name)
	if err != nil {
		return false, err
	}

	fm := fd.Mode()
	return fm.IsDir(), err
}
