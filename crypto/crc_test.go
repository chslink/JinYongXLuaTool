package crypto

import (
	"testing"
)

func TestCrc32(t *testing.T) {
	if GetCRC32([]byte("123456")) != 3639635087 {
		t.Fatal("crc32 error")
	}

}

func TestGetCRC32(t *testing.T) {
	text := `@` + newEncText
	if GetCRC32([]byte(text)) != 2673404822 {
		t.Fatal("crc error")
	}
}
