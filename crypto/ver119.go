package crypto

import (
	"log"
)

var ver119 = &VersionConfig{
	Version:    "119",
	Key:        "4dd83c28e46e3998249465f3",
	IV:         "4dd583c628e746e8",
	FixXml:     true,
	FixXmlFunc: fix119XmlFunc,
}

func fix119XmlFunc(filename string, data []byte) ([]byte, error) {
	if isXml(data) {
		return data, nil
	}
	dataLen := len(data)
	pos := 0

	for pos > 0 && pos < 200 && pos < dataLen {
		fixBit := pos % 3
		switch fixBit {
		case 0:
			data[pos] = data[pos] + 1
		case 1:
			data[pos] = data[pos] - 1
		case 2:
			data[pos] = data[pos] + 2
		}
		pos++
	}
	if pos > 0 {
		log.Printf("修复乱码... \n")
	}
	return data, nil
}
