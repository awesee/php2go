package php

import (
	"fmt"
	"log"
	"testing"
)

func TestFilePutContents(t *testing.T) {
	data := "this is a test string"
	dataByte := []byte(data)
	err := FilePutContents("testdata/test_fpc.txt", dataByte)
	if err != nil {
		t.Error(err)
		log.Fatal(err)
	}
}

func TestArrayKeys(t *testing.T) {
	var originData = map[string]interface{}{"foo": 123, "bar": "abc"}
	res := ArrayKeys(originData)
	fmt.Println(res)
}
