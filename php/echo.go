package php

//输出一个或多个字符串
func Echo(args ...string) {
	for _, v := range args {
		print(v)
	}
}
