package main

import (
	"fmt"

	"github.com/openset/php2go/php"
)

func main() {

	php.Echo("Hello ", "world!\n")

	php.Print("This is a string.\n")

	fmt.Println(php.Chr(65))

	fmt.Println(php.Ord('A'))

	fmt.Println(php.Array(1, 'a', "ABC"))

	fmt.Println(php.Strrev("Hello world!"))

	fmt.Println(php.Strrev("你好，世界"))

	fmt.Println(php.ArrayReverse(php.Array(1, 'a', "ABC")))

	fmt.Println(php.SysGetTempDir())

	fmt.Println(php.IsNumeric("-123.45"))

	fmt.Println(php.Basename("abc/cde.ext", "xt"))

	fmt.Println(php.Addcslashes("abc/cde.ext", 'c'))

	fmt.Println(php.Addslashes("abc/'\"c\\de.ext"))

	fmt.Println(php.Htmlspecialchars("This is some <b>bold</b> text."))

	fmt.Println(php.MbStrlen("你好，世界"))

	fmt.Println(php.Intval("-123abc"))

}
