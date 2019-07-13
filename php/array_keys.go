package php

// ArrayKeys - get keys of map data as a Array
// in php,the keys you want always is string or number
// here,let it be string
func ArrayKeys(data map[string]interface{}) []string {
	if len(data) < 1 {
		return []string{}
	}
	var resData []string
	for index := range data {
		resData = append(resData, index)
	}
	return resData
}
