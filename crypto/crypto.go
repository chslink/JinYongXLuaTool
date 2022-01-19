package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"

    "github.com/pkg/errors"
)

var key = []byte(`64555156d26643a06b691e44`)
var iv = []byte(`6451551f56dc266a`)

func Decrypt(data []byte) ([]byte, error) {
    enc := base64.StdEncoding
    dbuf := make([]byte, enc.DecodedLen(len(data)))
    n, err := enc.Decode(dbuf, data)
    if err != nil {
        return nil, errors.Wrap(err, "jxcrypt:decode base64 failed")
    }
    dbuf = dbuf[:n]
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, errors.Wrap(err, "jxcrypt:decrypt failed")
    }
    bm := cipher.NewCBCDecrypter(block, iv)
    bm.CryptBlocks(dbuf, dbuf)
    deData, err := NewPkcs7Padding(block.BlockSize()).Unpad(dbuf)
    if err != nil {
        return nil, errors.Wrap(err, "jxcrypt:decrypt failed")
    }
    return deData, nil
}

func Encrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, errors.Wrap(err, "jxcrypt:encrypt failed")
    }

    pbuf, err := NewPkcs7Padding(block.BlockSize()).Pad(data)
    if err != nil {
        return nil, errors.Wrap(err, "jxcrypt:encrypt failed")
    }
    bm := cipher.NewCBCEncrypter(block, iv)
    bm.CryptBlocks(pbuf, pbuf)

    buf := make([]byte, base64.StdEncoding.EncodedLen(len(pbuf)))
    base64.StdEncoding.Encode(buf, pbuf)
    return buf, nil
}
