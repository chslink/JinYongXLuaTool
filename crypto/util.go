package crypto

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// DecryptFile 解密文件
func DecryptFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("jxcrypt:open file failed:%s", file))
	}
	enText, _, err := splitFile(data)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("jxcrypt:open file failed:%s", file))
	}
	return Decrypt(enText)
}

func splitFile(data []byte) (context, sum []byte, err error) {
	split := bytes.Split(data, []byte("#"))
	if len(split) >= 2 {
		return split[0][1:], split[1], nil
	}
	return nil, nil, errors.New("jxcrypt: split context error")
}

// EncryptFile 加密文件
func EncryptFile(filename string, data []byte) error {
	enData, err := Encrypt(data)
	if err != nil {
		return errors.Wrap(err, "jxcrypt:encrypt failed")
	}
	enData = append([]byte("@"), enData...)
	enData = append(enData, []byte("#1671935116")...)
	return ioutil.WriteFile(filename, enData, os.ModePerm)
}

// EncryptDir 加密文件夹
func EncryptDir(src, dest string, ext ...string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !hasExt(path, ext) {
			log.Printf("...跳过不支持文件:[%s]\n", path)
			return nil
		}
		log.Printf("操作文件...[%s]\n", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		return EncryptFile(filepath.Join(dest, relPath), data)
	})
}

// DecryptDir 解密文件夹
func DecryptDir(src, dest string, ext ...string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !hasExt(path, ext) {
			log.Printf("...跳过不支持文件:[%s]\n", path)
			return nil
		}
		log.Printf("操作文件...[%s]\n", path)
		data, err := DecryptFile(path)
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filepath.Join(dest, relPath), data, os.ModePerm)
	})
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
