package php

import "strings"

type ArrayMap map[string]interface{}

const (
	CaseLOWER = iota
	CaseUPPER
)

//Changes the case of all keys in an array
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
