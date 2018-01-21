package php

import "time"

//Return current Unix timestamp
func Time() int64 {
	return time.Now().Unix()
}
