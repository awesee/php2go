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
}
