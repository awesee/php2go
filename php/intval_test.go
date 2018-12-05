package php

import "testing"

func TestIntval(t *testing.T) {

	m := map[string]int{
		"":         0,
		"0":        0,
		"-1":       -1,
		"+123abc":  123,
		"-123 abc": -123,
	}

	for str, i := range m {
		if r, _ := Intval(str); r != i {
			t.Errorf("Intval() should have returned [%d] for string [%s]", i, str)
		}
	}
}
