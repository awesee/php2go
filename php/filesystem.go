package php

import "os"

//Changes file mode
func Chmod(name string, mode os.FileMode) error {

	return os.Chmod(name, mode)
}

// +------------------------------------------------------------
// | @desc      Chown changes the numeric uid and gid of the named file.
// | @param     name string, uid int, gid int
// | @return    error
// |
// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Chown(name string, uid int, gid int) error {

	return os.Chown(name, uid, gid)
}

// +------------------------------------------------------------
// | @desc      make directort
// | @param     name string, mode os.FileMode
// | @return    error
// |
// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Mkdir(name string, mode os.FileMode) error {

	return os.Mkdir(name, mode)
}

// +------------------------------------------------------------
// | @desc      is directort
// | @param     name string
// | @return    bool , error
// |
// | @author    Openset <jinheking@sina.com>
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
