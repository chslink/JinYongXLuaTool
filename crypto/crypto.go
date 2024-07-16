package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"path/filepath"

	"github.com/pkg/errors"
)

func Decrypt(conf *VersionConfig, filename string, data []byte) ([]byte, error) {
	enc := base64.StdEncoding
	dbf := make([]byte, enc.DecodedLen(len(data)))
	n, err := enc.Decode(dbf, data)
	if err != nil {
		return nil, errors.Wrap(err, "decode base64 failed")
	}
	dbf = dbf[:n]
	fileType := filepath.Ext(filename)
	curKey := []byte(conf.Key)
	curIV := []byte(conf.IV)
	switch fileType {
	case ".lua":
		if conf.LuaKey != "" {
			curKey = []byte(conf.LuaKey)
			curIV = []byte(conf.LuaIV)
		}
	case ".xml":
		if conf.XmlKey != "" {
			curKey = []byte(conf.XmlKey)
			curIV = []byte(conf.XmlIV)
		}
	}

	block, err2 := aes.NewCipher(curKey)
	if err2 != nil {
		return nil, errors.Wrap(err, "decrypt failed")
	}
	bm := cipher.NewCBCDecrypter(block, curIV)
	bm.CryptBlocks(dbf, dbf)
	deData, err := NewPkcs7Padding(block.BlockSize()).Unpad(dbf)
	if err != nil {
		log.Printf("warning: unpad file failed: %v,%s", err, filename)
		deData = dbf
	}
	switch fileType {
	case ".lua":
		if conf.FixLua {
			return conf.FixLuaFunc(filename, deData)
		}
	case ".xml":
		if conf.FixXml {
			return conf.FixXmlFunc(filename, deData)
		}

	}
	return deData, nil
}

func Encrypt(conf *VersionConfig, filename string, data []byte) ([]byte, error) {
	fileType := filepath.Ext(filename)
	curKey := []byte(conf.Key)
	curIV := []byte(conf.IV)
	switch fileType {
	case ".lua":
		if conf.LuaKey != "" {
			curKey = []byte(conf.LuaKey)
			curIV = []byte(conf.LuaIV)
		}
	case ".xml":
		if conf.XmlKey != "" {
			curKey = []byte(conf.XmlKey)
			curIV = []byte(conf.XmlIV)
		}
	}
	block, err := aes.NewCipher(curKey)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt failed")
	}

	pbuf, err := NewPkcs7Padding(block.BlockSize()).Pad(data)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt failed")
	}
	bm := cipher.NewCBCEncrypter(block, curIV)
	bm.CryptBlocks(pbuf, pbuf)

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(pbuf)))
	base64.StdEncoding.Encode(buf, pbuf)
	return buf, nil
}
