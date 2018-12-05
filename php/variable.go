package php

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// Boolval - Get the boolean value of a variable
func Boolval(val interface{}) bool {

	switch v := val.(type) {
	case int, int8, int16, int32, int64:
		if v != 0 {
			return true
		}
		return false
	case uint, uint8, uint16, uint32, uint64:
		if v != 0 {
			return true
		}
		return false
	case bool:
		return v
	case complex64, complex128:
		if v != complex128(0) {
			return true
		}
		return false
	case float32, float64:
		if v != float64(0) {
			return true
		}
	default:
		return v == nil
	}

	return false
}

// Empty - Determine whether a variable is empty
func Empty(v interface{}) bool {

	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Bool:
		return !val.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return val.Float() == 0.00

	case reflect.Complex64, reflect.Complex128:
		return val.Complex() == 0+0i

	case reflect.String:
		realVal := val.String()
		return realVal == "" || realVal == "0"

	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return val.Len() == 0

	case reflect.Struct:
		return val.NumField() == 0
	}

	return false
}

// Intval - Get the integer value of a variable
func Intval(str string) (int, error) {

	str = strings.TrimSpace(str)
	pre := ""
	if strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+") {
		pre = str[0:1]
		str = str[1:]
	}

	i := strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	if i > -1 {
		str = str[0:i]
	}
	if str == "" {
		str = "0"
	}

	return strconv.Atoi(pre + str)
}

// Strval - Get string value of a variable
func Strval(val interface{}) string {

	return fmt.Sprintf("%v", val)
}

// Gettype - Get the type of a variable
func Gettype(v interface{}) string {

	// t := reflect.TypeOf(v)
	//
	// return t.String()
	return fmt.Sprintf("%T", v)
}

// IsBool - Finds out whether a variable is a boolean
func IsBool(v interface{}) bool {

	// _, ok := v.(bool)
	//
	// return ok
	return v == true || v == false
}
