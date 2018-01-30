package php

import (
	"io"
	"io/ioutil"
	"os"
)

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

// +------------------------------------------------------------
// | @desc     copy file
// | @param     dstName string, srcName string
// | @return    int64 , error

// | @since     https://studygolang.com/articles/1599

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Copy(dstName string, srcName string) (written int64, err error) {

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// +------------------------------------------------------------
// | @desc      close file
// | @param     file *os.File
// | @return    error

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Fclose(file *os.File) error {

	return file.Close()
}

// +------------------------------------------------------------
// | @desc      ReadDir reads the directory named by dirname and returns a list of directory entries sorted by filename.
// | @param     dirPth string
// | @return    []os.FileInfo, error

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/26
// +------------------------------------------------------------
func Dirname(dirPth string) ([]os.FileInfo, error) {

	return ioutil.ReadDir(dirPth)
}

// +------------------------------------------------------------
// | @desc      Remove file.
// | @param     file string
// | @return     error

// | @author    Openset <jinheking@sina.com>
// | @link      https://github.com/sunnyregion
// | @date      2018/01/30
// +------------------------------------------------------------
func Delete(file string) error {

	return os.Remove(file)
}
