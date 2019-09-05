package php

import "testing"

// TestFilePutContents
func TestFilePutContents(t *testing.T) {
	input := []byte("this is a test string")
	err := FilePutContents("testdata/test_fpc.txt", input)
	if err != nil {
		t.Fatalf("output: %v, expected: %v", err, nil)
	}
}
