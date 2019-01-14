//
//  findHost.go
//  WTest
//
//  Created by 吴道睿 on 15/12/7.
//  Copyright (c) 2015年 吴道睿. All rights reserved.
//

package discovery

import (
	"log"
	"net"
	"net/http"
	"strings"
)

/*
 找到一个可用的端口
*/
func FindUseAbleHost() (host string) {
	// 获取本机ip
	conn, err := net.Dial("udp", "www.tongdun.cn:80")
	if err != nil {
		log.Println("GetLocalIpErr:", err)
		return
	}
	conn.Close()

	/*
	 这里或许可以直接使用 conn.LocalAddr().String() 作为返回用来开启服务
	 需要测试
	*/
	loaclIp := strings.Split(conn.LocalAddr().String(), ":")[0]

	// 随机选择可用端口
	conn_1, err_1 := net.Listen("tcp", "127.0.0.1:0")
	if err_1 != nil {
		log.Println("GetLocalPortErr:", err_1)
		return
	}
	_, port, _ := net.SplitHostPort(conn_1.Addr().String())
	conn_1.Close()
	host = loaclIp + ":" + port
	return
}

func IsHttpUsablePort(addr, port string) bool {
	resp, err := http.Get("http://" + addr + ":" + port)
	if err != nil {
		return true
	}
	defer resp.Body.Close()
	return false
}
