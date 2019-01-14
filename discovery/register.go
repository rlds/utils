//
//  register.go
//  WTest
//
//  Created by 吴道睿 on 15/12/4.
//  Copyright (c) 2015年 吴道睿. All rights reserved.
//
package discovery

import (
	"net"
	"net/rpc"
	"rlds/rlog"
	"time"
)

type Register struct {
	registerHostList []string
	rpcClient        *rpc.Client
	stepRegTime      time.Duration
	workerInfo       WorkerInfo
}

type WorkerInfo struct {
	Prekey string
	Host   string
	Type   string
	Name   string
}

/*
   开启注册并启动心跳连接
*/
func NewRegister(preType, wokerType, workerName, host string, registerHostList []string) (r *Register) {
	r = new(Register)
	r.registerHostList = registerHostList
	r.workerInfo.Prekey = preType
	r.workerInfo.Type = wokerType
	r.workerInfo.Name = workerName
	r.workerInfo.Host = host
	r.stepRegTime = time.Second * 2
	go r.doKeep()
	return
}

//中途需要重新设置服务信息的时候执行
func (r *Register) ReSetWorkerInfo(preType, wokerType, workerName, host string) {
	r.workerInfo.Prekey = preType
	r.workerInfo.Type = wokerType
	r.workerInfo.Name = workerName
	r.workerInfo.Host = host
}

func conToServer(host string) (rpcClient *rpc.Client, ok bool) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		rlog.V(1).Info("tcp Dial err:", err)
		return
	}
	addr := conn.LocalAddr()
	rlog.V(1).Info("discovery_rpc_conn[" + host + "] at:" + addr.String())
	rpcClient = rpc.NewClient(conn)
	ok = rpcClient != nil
	return
}

func (r *Register) getOneRpcClient() (ok bool) {
	for _, host := range r.registerHostList {
		rlog.V(1).Info("reg test con:", host)
		r.rpcClient, ok = conToServer(host)
		if ok {
			break
		}
	}
	return
}

func (r *Register) doKeep() {
	ok := r.getOneRpcClient()
	var err error
	for {
		if r.rpcClient == nil {
			ok = r.getOneRpcClient()
			if !ok {
				rlog.V(1).Info("no disServer !!")
				goto EndSleep
			}
		}

		//执行注册
		err = r.rpcClient.Call("Drpc.Register", r.workerInfo, &ok)
		if err != nil || !ok {
			rlog.V(1).Info("call Register err:", err)
			r.rpcClient = nil
		}
	EndSleep:
		time.Sleep(r.stepRegTime)
	}
}
