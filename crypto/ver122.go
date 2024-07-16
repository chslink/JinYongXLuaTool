package crypto

import (
	"log"
	"strings"
)

var ver122 = &VersionConfig{
	Version:    "122",
	Key:        "88672fb343269ae765e96977",
	IV:         "886e72f5b341326e",
	FixXml:     true,
	FixXmlFunc: fix122xml,
	FixLua:     true,
	FixLuaFunc: fix122lua,
}

func fix122lua(filename string, data []byte) ([]byte, error) {
	if strings.HasSuffix(filename, "EncodingLib.lua") {
		return data, nil
	}
	pos := 0
	dataLen := len(data)
	for pos < 200 && pos < dataLen {
		fixBit := pos % 2
		switch fixBit {
		case 0:
			data[pos] = data[pos] - 3
		case 1:
			data[pos] = data[pos] - 4
		}
		pos++
	}
	if pos > 0 {
		log.Printf("修复乱码... \n")
	}
	return data, nil
}

func fix122xml(filename string, data []byte) ([]byte, error) {
	if isXml(data) {
		return data, nil
	}
	pos := 0
	dataLen := len(data)
	for pos >= 0 && pos < 60 && pos < dataLen {
		fixBit := pos % 2
		switch fixBit {
		case 0:
			data[pos] = data[pos] + 1
		case 1:
			data[pos] = data[pos] + 3
		}
		pos++
	}
	if pos > 0 {
		log.Printf("修复乱码... \n")
	}
	return data, nil
}
