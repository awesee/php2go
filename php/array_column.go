package php

// ArrayColumn — Return the values from a single column in the input array
func ArrayColumn(arrayMap map[string]map[string]interface{}, columnKey string) (r []interface{}) {

	for _, input := range arrayMap {
		if v, ok := input[columnKey]; ok {
			r = append(r, v)
		}
	}

	return
}
