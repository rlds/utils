package utils

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
)

// ByteToUint64 ByteToUint64
func ByteToUint64(b []byte) (r uint64) {
	r = (uint64)(b[7])
	r |= (uint64)(b[6]) << 8
	r |= (uint64)(b[5]) << 16
	r |= (uint64)(b[4]) << 24
	r |= (uint64)(b[3]) << 32
	r |= (uint64)(b[2]) << 40
	r |= (uint64)(b[1]) << 48
	r |= (uint64)(b[0]) << 56
	return
}

// Uint64ToByte Uint64ToByte
func Uint64ToByte(u uint64) (b []byte) {
	b = make([]byte, 8)
	b[7] = (byte)(u)
	b[6] = (byte)(u >> 8)
	b[5] = (byte)(u >> 16)
	b[4] = (byte)(u >> 24)
	b[3] = (byte)(u >> 32)
	b[2] = (byte)(u >> 40)
	b[1] = (byte)(u >> 48)
	b[0] = (byte)(u >> 56)
	return
}

// BytesToUint64  Converts bytes to an integer
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Uint64ToBytes  Converts a uint to a byte slice
func Uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}

// StrToBytes Converts string to a byte slice
func StrToBytes(s string) []byte {
	return []byte(s)
}

//BytesToStr  Converts bytes to string
func BytesToStr(b []byte) string {
	return string(b)
}

//IntToStr  Converts int to string
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// StrToInt Converts string to int
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// Uint64ToStr Converts uint64 to string
func Uint64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// StrToUint64 Converts string to uint64
func StrToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// UnicodeToUtf8 Converts unicode to utf-8
func UnicodeToUtf8(strUnicode string) string {
	unilen := 6
	current, step := 0, 0
	ret := ""
	for {
		if current >= len(strUnicode) {
			break
		}
		start := strings.Index(strUnicode[current:], `\u`)
		if start == -1 {
			ret += strUnicode[current:]
			break
		}
		if tmp, err := unicode2Str(strUnicode[current+start : current+start+unilen]); err == nil {
			ret += strUnicode[current:current+start] + tmp
			step = start + unilen
		} else {
			step = 2
		}
		current += step
	}
	return ret
}

func unicode2Str(strUnicode string) (string, error) {
	ret := strings.Replace(strUnicode, "\\u", "", -1)
	tmp, err := hex.DecodeString(ret)
	if err != nil {
		return "", err
	}
	dec := mahonia.NewDecoder("utf-16")
	return dec.ConvertString(string(tmp)), nil
}
