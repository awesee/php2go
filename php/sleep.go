package php

import "time"

// Sleep - Delay execution
func Sleep(s int64) {
	time.Sleep(time.Duration(s) * time.Second)
}

// Usleep - Delay execution in microseconds
func Usleep(ms int64) {
	time.Sleep(time.Duration(ms) * time.Microsecond)
}
