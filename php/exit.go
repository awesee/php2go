package php

import "os"

// Exit — Output a message and terminate the current script
func Exit(code int) {
	os.Exit(code)
}

// Die — Equivalent to exit
func Die(code int) {
	Exit(code)
}
