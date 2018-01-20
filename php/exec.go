package php

import "os/exec"

//Execute an external program
func Exec(s string) {
	exec.Command(s).Run()
}
