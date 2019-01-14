package utils

import "strings"

func GetIpPort(host string) (ip, port string) {
	arr := strings.Split(host, ":")
	if len(arr) == 2 {
		ip, port = arr[0], arr[1]
	}
	return
}
