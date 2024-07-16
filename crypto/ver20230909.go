package crypto

import (
	"log"
)

var ver20230909 = &VersionConfig{
	Version:    "20230909",
	LuaKey:     "mjamjqzjimtiowiyjkoweytu",
	LuaIV:      "mja5mjqzzjiymti5",
	FixLua:     true,
	FixLuaFunc: fix122lua,
	XmlKey:     "1/832e1018a28888`4625`c4",
	XmlIV:      "1/8132e110188a2a",
	FixXml:     true,
	FixXmlFunc: fix20230909xml,
}

func fix20230909xml(filename string, data []byte) ([]byte, error) {
	if isXml(data) {
		return data, nil
	}
	pos := 0
	dataLen := len(data)
	for pos >= 0 && pos < 101 && pos < dataLen {
		fixBit := pos % 2
		switch fixBit {
		case 0:
			data[pos] = data[pos] + 1
		case 1:
			data[pos] = data[pos] + 2
		}
		pos++
	}
	if pos > 0 {
		log.Printf("修复乱码... \n")
	}
	return data, nil
}
