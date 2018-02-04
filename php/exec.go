package php

import "os/exec"

// Exec - Execute an external program
func Exec(s string) {

	exec.Command(s).Run()
}
