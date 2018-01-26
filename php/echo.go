package php

import "fmt"

//Output one or more strings
func Echo(args ...interface{}) {

	fmt.Print(args...)
}
