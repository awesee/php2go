package php

import "os/exec"

//执行一个外部程序
func Exec(s string) {
	exec.Command(s).Run()
}
