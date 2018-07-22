package php

import (
	"testing"
)

// +------------------------------------------------------------
// | @desc      Changes the case of all keys in an array
// | @param     arr  ArrayMap
// | @return    ArrayMap
// |
// | @author    Openset <openset.wang@gmail.com>
// | @link      https://github.com/openset
// | @date      2018/01/24
// +------------------------------------------------------------
func TestMbStrlen(t *testing.T) {

	s := "A"
	if MbStrlen(s) != 1 {
		t.Errorf("MbStrlen() should have returned 1 for string [%s]", s)
	}

	s = "abc"
	if MbStrlen(s) != 3 {
		t.Errorf("MbStrlen() should have returned 3 for string [%s]", s)
	}

	s = "你好，世界"
	if MbStrlen(s) != 5 {
		t.Errorf("MbStrlen() should have returned 5 for string [%s]", s)
	}
}
