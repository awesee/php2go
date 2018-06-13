package php_test

import (
	"testing"

	. "github.com/openset/php2go/php"
)

func TestChr(t *testing.T) {

	x := 65
	if Chr(x) != "A" {
		t.Errorf("Chr() should have returned A for [%d]", x)
	}

	x += 256
	if Chr(x) != "A" {
		t.Errorf("Chr() should have returned A for [%d]", x)
	}

}
