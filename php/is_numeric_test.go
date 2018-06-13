package php_test

import (
	"testing"

	. "github.com/openset/php2go/php"
)

func TestIsNumericUint8(t *testing.T) {

	var x uint8

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint8 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint8 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint8 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint8 [%v]", x)
	}

	x = 255
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint8 [%v]", x)
	}
}

func TestIsNumericUint16(t *testing.T) {

	var x uint16

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint16 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint16 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint16 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint16 [%v]", x)
	}

	x = 255
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint16 [%v]", x)
	}
}

func TestIsNumericUint32(t *testing.T) {

	var x uint32

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint32 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint32 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint32 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint32 [%v]", x)
	}

	x = 255
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint32 [%v]", x)
	}
}

func TestIsNumericUint64(t *testing.T) {

	var x uint64

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint64 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint64 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint64 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint64 [%v]", x)
	}

	x = 255
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint64 [%v]", x)
	}
}

func TestIsNumericInt8(t *testing.T) {

	var x int8

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = -5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = -1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = -123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = 127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}

	x = -127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int8 [%v]", x)
	}
}

func TestIsNumericInt16(t *testing.T) {

	var x int16

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = -5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = -1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = -123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = 127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}

	x = -127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int16 [%v]", x)
	}
}

func TestIsNumericInt32(t *testing.T) {

	var x int32

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = -5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = -1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = -123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = 127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}

	x = -127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int32 [%v]", x)
	}
}

func TestIsNumericInt64(t *testing.T) {

	var x int64

	x = 0

	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}
	x = 5

	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = -5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = -1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = -123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = 127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}

	x = -127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int64 [%v]", x)
	}
}

func TestIsNumericFloat32(t *testing.T) {

	var x float32

	x = 0.0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float32 [%v]", x)
	}

	x = 1.0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float32 [%v]", x)
	}

	x = -1.0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float32 [%v]", x)
	}

	x = 12.34
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float32 [%v]", x)
	}

	x = -12.34
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float32 [%v]", x)
	}
}

func TestIsNumericFloat64(t *testing.T) {

	var x float64

	x = 0.0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float64 [%v]", x)
	}

	x = 1.0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float64 [%v]", x)
	}

	x = -1.0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float64 [%v]", x)
	}

	x = 12.34
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float64 [%v]", x)
	}

	x = -12.34
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for float64 [%v]", x)
	}
}

func TestIsNumericComplex64(t *testing.T) {

	var x complex64

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}

	x = 1i
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}

	x = -1i
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}

	x = 12.34i
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}

	x = -12.34i
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}

	x = 12.34 + 56.78i
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}

	x = -12.34 + -56.78i
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for complex64 [%v]", x)
	}
}

func TestIsNumericUint(t *testing.T) {

	var x uint

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint [%v]", x)
	}

	x = 255
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for uint [%v]", x)
	}
}

func TestIsNumericInt(t *testing.T) {

	var x int

	x = 0
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = 5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = -5
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = 1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = -1
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = 123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = -123
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = 127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}

	x = -127
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for int [%v]", x)
	}
}

func TestIsNumericString(t *testing.T) {

	var x string

	x = "0"
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for string [%v]", x)
	}

	x = "123"
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for string [%v]", x)
	}

	x = "-123"
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for string [%v]", x)
	}

	x = "123.45"
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for string [%v]", x)
	}

	x = "-123.45"
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned true for string [%v]", x)
	}

	x = "apple"
	if IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned false for string [%v]", x)
	}

	x = "1-23"
	if IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned false for string [%v]", x)
	}

	x = "123-"
	if IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned false for string [%v]", x)
	}

	x = ".123"
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned false for string [%v]", x)
	}

	x = "123."
	if !IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned false for string [%v]", x)
	}

	x = "1.2.3"
	if IsNumeric(x) {
		t.Errorf("IsNumeric() should have returned false for string [%v]", x)
	}
}
