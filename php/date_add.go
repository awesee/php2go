package php

import "time"

// DateAdd -  Adds an amount of days, months, years, hours, minutes and seconds to a DateTime object
func DateAdd(t time.Time, years int, months int, days int) time.Time {
	return t.AddDate(years, months, days)
}
