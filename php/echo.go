package php

//Output one or more strings
func Echo(args ...string) {
	for _, v := range args {
		print(v)
	}
}
