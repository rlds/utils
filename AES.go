package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

func safeBase64Encode(bt []byte) string {
	str := base64.StdEncoding.EncodeToString(bt)
	str = strings.Replace(str, "=", "", -1)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	return str
}

func safeBase64Decode(str string) ([]byte, error) {
	switch len(str) % 4 {
	case 0:
	case 2:
		str = str + "=="
	case 3:
		str = str + "="
	default:
		return nil, errors.New("illegal base64 Encode String")
	}
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)

	return base64.StdEncoding.DecodeString(str)
}

//AesEncrypt 加密
func AesEncrypt(key, plantText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plantText = PKCS7Padding(plantText, block.BlockSize())
	blockSize := block.BlockSize()
	blockModel := cipher.NewCBCEncrypter(block, key[:blockSize])

	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)

	ecsrt := safeBase64Encode(ciphertext)
	return []byte(ecsrt), nil
}

//AesDecrypt 解密
func AesDecrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	dst, errr := safeBase64Decode(string(ciphertext))
	if errr != nil {
		return nil, errr
	}

	if (len(dst) % blockSize) == 0 {
		blockModel := cipher.NewCBCDecrypter(block, key[:blockSize])

		plantText := make([]byte, len(dst))
		blockModel.CryptBlocks(plantText, dst)
		plantText = PKCS7UnPadding(plantText, block.BlockSize())
		return plantText, nil
	}
	return nil, errors.New("input not full blocks")
}

// PKCS7Padding PKCS7Padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(Get_rand(125))}, padding-1)
	padtext = append(padtext, byte(padding))
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding PKCS7UnPadding
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

/*
获取随机数
*/
func Get_rand(mx int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(mx)
}
