package main

import (
	"fmt"

	"github.com/sunnyregion/php2go/php"
)

func main() {

	//	php.Echo("Hello ", "world!\n")

	//	php.Print("This is a string.\n")

	//	println(php.Chr(65))

	//	println(php.Ord('A'))

	//	fmt.Println(php.Array(1, 'a', "ABC"))

	//	println(php.Strrev("Hello world!"))

	//	println(php.Strrev("你好，世界"))

	//	fmt.Println(php.ArrayReverse(php.Array(1, 'a', "ABC")))

	//	println(php.SysGetTempDir())

	//	println(php.IsNumeric("-123.45"))

	//	println(php.Basename("abc/cde.ext", "xt"))

	//	println(php.Addcslashes("abc/cde.ext", 'c'))

	//	println(php.Addslashes("abc/'\"c\\de.ext"))

	//	println(php.Htmlspecialchars("This is some <b>bold</b> text."))

	//	println(php.MbStrlen("你好，世界"))

	dir := php.Getcwd()
	fmt.Println(dir)

	b, err := php.IsDir(dir)

	if err != nil {
		println(err)
	} else {
		fmt.Println(b)
	}

}
