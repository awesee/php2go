package php

import "strings"

// ArrayMap is one of Array type
type ArrayMap map[string]interface{}

// Constants for ArrayChangeKeyCase
const (
	CaseLOWER = iota
	CaseUPPER
)

// ArrayChangeKeyCase - Changes the case of all keys in an array
func ArrayChangeKeyCase(arr ArrayMap, Case int) ArrayMap {

	var tmp = ArrayMap{}
	for k, v := range arr {
		if Case == CaseUPPER {
			tmp[strings.ToUpper(k)] = v
		} else if Case == CaseLOWER {
			tmp[strings.ToLower(k)] = v
		}
	}

	return tmp
}
