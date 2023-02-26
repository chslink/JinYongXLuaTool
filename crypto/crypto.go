package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/pkg/errors"
)

// var key = []byte(`64555156d26643a06b691e44`)
// var iv = []byte(`6451551f56dc266a`)
var key = []byte(`4dd83c28e46e3998249465f3`)
var iv = []byte(`4dd583c628e746e8`)

var keyM = map[string]KeyPair{
	"111": {
		Key: "64555156d26643a06b691e44",
		IV:  "6451551f56dc266a",
	},
	"119": {
		Key: "4dd83c28e46e3998249465f3",
		IV:  "4dd583c628e746e8",
	},
}

type KeyPair struct {
	Key string
	IV  string
}

func SetKeyVer(ver string) {
	if v, ok := keyM[ver]; ok {
		key = []byte(v.Key)
		iv = []byte(v.IV)
	}
}
func SetKeyIV(_key, _iv string) {
	key = []byte(_key)
	iv = []byte(_iv)
}

func Decrypt(data []byte) ([]byte, error) {
	enc := base64.StdEncoding
	dbf := make([]byte, enc.DecodedLen(len(data)))
	n, err := enc.Decode(dbf, data)
	if err != nil {
		return nil, errors.Wrap(err, "decode base64 failed")
	}
	dbf = dbf[:n]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "decrypt failed")
	}
	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(dbf, dbf)
	deData, err := NewPkcs7Padding(block.BlockSize()).Unpad(dbf)
	if err != nil {
		return nil, errors.Wrap(err, "decrypt failed")
	}
	return deData, nil
}

func Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt failed")
	}

	pbuf, err := NewPkcs7Padding(block.BlockSize()).Pad(data)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt failed")
	}
	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(pbuf, pbuf)

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(pbuf)))
	base64.StdEncoding.Encode(buf, pbuf)
	return buf, nil
}
