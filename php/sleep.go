package php

import "time"

//Sleep - Delay execution
func Sleep(s int64) {

	time.Sleep(time.Duration(s) * time.Second)
}
