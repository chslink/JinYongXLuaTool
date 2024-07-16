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
	text := `@` + var122EncText
	if GetCRC32([]byte(text)) != 4101382441 {
		t.Fatal("crc error")
	}
}
