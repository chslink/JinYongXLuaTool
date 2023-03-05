package crypto

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	ModeDe = "d"
	ModeEn = "e"
)

// DecryptFile 解密文件
func DecryptFile(data []byte) ([]byte, error) {
	enText, _, err := splitFile(data)
	if err != nil {
		return nil, err
	}
	return Decrypt(enText)
}

func splitFile(data []byte) (context, sum []byte, err error) {
	split := bytes.Split(data, []byte("#"))
	if len(split) < 2 {
		return nil, nil, errors.Errorf("split context error:%s", string(data))
	}
	return split[0][1:], split[1], nil
}

// EncryptFile 加密文件
func EncryptFile(data []byte) ([]byte, error) {
	enData, err := Encrypt(data)
	if err != nil {
		return nil, err
	}
	enData = append([]byte("@"), enData...)
	enData = append(enData, []byte(fmt.Sprintf("#%d", GetCRC32(enData)))...)
	return enData, nil
}

func RunDir(src, dst, mode string, ext []string) (err error) {
	isEnc := false
	switch mode {
	case ModeEn:
		isEnc = false
	case ModeDe:
		isEnc = true
	default:
		detach := false
		err = walkDir(src, ext, func(filename string, b []byte) error {
			if detach {
				return filepath.SkipDir
			}
			isEnc = isEncrypt(b)
			detach = true
			return filepath.SkipDir
		})
	}
	if err != nil {
		return err
	}

	_ = os.MkdirAll(dst, os.ModeDir)
	if !isEnc {
		log.Printf("正在加密...")

		err = walkDir(src, ext, func(filename string, b []byte) error {
			if isEncrypt(b) {
				return writeFile(src, dst, filename, b)
			}
			data, err2 := EncryptFile(b)
			if err2 != nil {
				return err2
			}
			return writeFile(src, dst, filename, data)
		})
	} else {
		log.Printf("正在解密...")

		err = walkDir(src, ext, func(filename string, b []byte) error {
			if !isEncrypt(b) {
				return writeFile(src, dst, filename, b)
			}
			data, err2 := DecryptFile(b)
			if err2 != nil {
				return err2
			}
			// 119版本xml乱码修复
			if strings.Contains(filename, ".xml") && curVer == "119" {
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
			// 122 版本修复
			if curVer == "122" {
				if strings.Contains(filename, ".xml") {
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
				}
			}
			return writeFile(src, dst, filename, data)
		})
	}
	return err
}

func hasExt(path string, ext []string) bool {
	e := filepath.Ext(path)
	for _, v := range ext {
		if v == e {
			return true
		}
	}
	return false
}

func isEncrypt(data []byte) bool {
	return bytes.HasPrefix(data, []byte("@"))
}

func writeFile(src, dst, filename string, data []byte) error {
	relPath, err := filepath.Rel(src, filename)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dst, relPath), data, os.ModePerm)
}

func walkDir(dir string, ext []string, fun func(filename string, b []byte) error) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !hasExt(path, ext) {
			log.Printf("...跳过不支持文件:[%s]\n", path)
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		log.Printf("操作文件...[%s]\n", path)
		return fun(path, data)
	})
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
