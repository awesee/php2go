package main

import (
	"fmt"
	"github.com/openset/php2go/php"
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
}
