package test

import (
	"fmt"
	"log"
	"php2go/php"
	"testing"
)

func TestFilePutContents(t *testing.T) {
	data := "who are u"
	dataByte := []byte(data)
	err := php.FilePutContents("test/test_fpc.txt",dataByte)
	if err!=nil {
		t.Error(err)
		log.Fatal(err)
	}
}

func TestArrayKeys(t *testing.T) {
	var orginData = map[string]interface{}{"1":1,"nishi":12,}
	res := php.ArrayKeys(orginData)
	fmt.Println(res)
}
