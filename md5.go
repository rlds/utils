package utils

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"io"
)

/*
   md5相关
*/

// GetBytesToMd5Str 从字符串到md5
func GetBytesToMd5Str(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// GetStrMd5ToInt md5字符串 转为数字
func GetStrMd5ToInt(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// GetBytesMd5ToInt md5字符串 转为数字
func GetBytesMd5ToInt(key []byte) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

//GetStringMd5 对字符串进行MD5哈希
func GetStringMd5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return hex.EncodeToString(t.Sum(nil))
}
