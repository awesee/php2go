package main

import (
	"fmt"
	"./php"
)

func main() {
	php.Echo("Hello ", "world!\n")
	php.Print("This is a string.\n")
	println(php.Chr(65))
	println(php.Ord('A'))
	fmt.Println(php.Array(1, 'a', "ABC"))
	fmt.Println(php.Strrev("Hello world!"))
	fmt.Println(php.Strrev("你好，世界"))
	fmt.Println(php.ArrayReverse(php.Array(1, 'a', "ABC")))
	println(php.SysGetTempDir())
	println(php.IsNumeric("-123.45"))
	fmt.Println(php.Basename("abc/cde.ext", "xt"))
	fmt.Println(php.Addcslashes("abc/cde.ext", 'c'))
	fmt.Println(php.Addslashes("abc/'\"c\\de.ext"))

	println(php.Htmlspecialchars("This is some <b>bold</b> text."))
}
