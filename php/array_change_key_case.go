package php

import "strings"

//ArrayMap Array type
type ArrayMap map[string]interface{}

//const
const (
	CaseLOWER = iota
	CaseUPPER
)

// ArrayChangeKeyCase - Changes the case of all keys in an array
// +------------------------------------------------------------
// | @param     arr  ArrayMap
// | @return    ArrayMap
// |
// | @author    Openset <openset.wang@gmail.com>
// | @link      https://github.com/openset
// | @date      2018/01/24
// +------------------------------------------------------------
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
