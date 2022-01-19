package main

import (
	"fmt"
	"testing"
)

func TestCrc32(t *testing.T) {
	if GetCRC32([]byte("123456")) != 3639635087 {
		t.Fatal("crc32 error")
	}

}

func TestGetCRC32(t *testing.T) {
	text := `@` + newEncText
	fmt.Println(GetCRC32([]byte(text)))
}
