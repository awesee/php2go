package php

//将一个或多个单元压入数组的末尾（入栈）
func ArrayPush(s *[]string, args ...string) {
	for _, v := range args {
		*s = append(*s, v)
	}
}
