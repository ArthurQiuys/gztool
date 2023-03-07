package util

import (
	"log"
	"testing"
)

func TestBinaryExist(t *testing.T) {
	AppExist("sssss")
	if AppExist("sssss") {
		log.Fatal("sssss shouldn't exist")
	}
	if !AppExist("goctl") {
		log.Fatal("goctl should exist")
	}
}
