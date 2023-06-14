package crypto

import (
	"encoding/xml"
	"log"
	"strings"
)

type FixFunc func(filename string, data []byte)

var fixMap = map[string]FixFunc{
	"119": fix119Xml,
	"122": fix122,
}

func Fix(filename string, data []byte) {
	fun, ok := fixMap[curVer]
	if !ok {
		return
	}
	fun(filename, data)
}

func fix119Xml(filename string, data []byte) {
	if !strings.HasSuffix(filename, ".xml") {
		return
	}
	pos := 0
	for (pos > 0 && pos < 200) || !isXml(data) {
		fixBit := pos % 3
		switch fixBit {
		case 0:
			data[pos] = data[pos] + 1
		case 1:
			data[pos] = data[pos] - 1
		case 2:
			data[pos] = data[pos] + 2
		}
		pos += 1
	}
	if pos > 0 {
		log.Printf("修复乱码... \n")
	}
}

func fix122(filename string, data []byte) {
	if strings.HasSuffix(filename, ".xml") {
		pos := 0
		for (pos > 0 && pos < 60) || !isXml(data) {
			fixBit := pos % 2
			switch fixBit {
			case 0:
				data[pos] = data[pos] + 1
			case 1:
				data[pos] = data[pos] + 3
			}
			pos += 1
		}
		if pos > 0 {
			log.Printf("修复乱码... \n")
		}
		return
	}
	if strings.HasSuffix(filename, ".lua") {
		if strings.HasSuffix(filename, "EncodingLib.lua") {
			return
		}
		pos := 0
		for pos < 200 && pos < len(data) {
			fixBit := pos % 2
			switch fixBit {
			case 0:
				data[pos] = data[pos] - 3
			case 1:
				data[pos] = data[pos] - 4
			}
			pos += 1
		}
		if pos > 0 {
			log.Printf("修复乱码... \n")
		}
	}

}

func isXml(data []byte) bool {
	err := xml.Unmarshal(data, new(interface{}))
	if err == nil {
		return true
	}
	//if strings.Contains(err.Error(), "attribute name without = in element") {
	//	return true
	//}

	return false
}

func init() {

}
