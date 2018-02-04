package php

import "fmt"

// Echo - Output one or more strings
func Echo(args ...interface{}) {

	fmt.Print(args...)
}
