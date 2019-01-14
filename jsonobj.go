package utils

import (
	"encoding/json"
)

// ObjToBytes ObjToBytes
func ObjToBytes(da interface{}) []byte {
	data, _ := json.Marshal(da)
	return data
}

// ObjToStr ObjToStr
func ObjToStr(da interface{}) string {
	data, _ := json.Marshal(da)
	return string(data)
}
