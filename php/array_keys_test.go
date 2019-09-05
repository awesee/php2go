package php

import (
	"reflect"
	"sort"
	"testing"
)

// TestArrayKeys
func TestArrayKeys(t *testing.T) {
	input := map[string]interface{}{"foo": 123, "bar": "abc"}
	expected := []string{"bar", "foo"}
	output := ArrayKeys(input)
	sort.Strings(output)
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("output: %v, expected: %v", output, expected)
	}
}
