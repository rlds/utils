//
//  discovery.go
//  WTest
//
//  Created by 吴道睿 on 15/12/4.
//  Copyright (c) 2015年 吴道睿. All rights reserved.
//
package discovery

import (
	"net/rpc"
	"rlds/rlog"
)

type Discovery struct {
	disServerHostList []string
	rpcClient         *rpc.Client
	workerInfo        WorkerInfo
}

func NewDiscovery(preType, wokerType, workerName string, registerHostList []string) (dis *Discovery) {
	dis = new(Discovery)
	dis.disServerHostList = registerHostList
	dis.workerInfo.Prekey = preType
	dis.workerInfo.Type = wokerType
	dis.workerInfo.Name = workerName
	return
}

func (d *Discovery) getOneRpcClient() (ok bool) {
	for _, host := range d.disServerHostList {
		rlog.V(3).Info("dis_get:", host)
		d.rpcClient, ok = conToServer(host)
		if ok {
			break
		}
	}
	return
}

type Wret struct {
	Ok   bool
	Host string
}

/*
 返回数组
*/
type WlistRet struct {
	Num   int
	HList []string
}

//根据类型
func (d *Discovery) GetWorkerByTypeAndName(workerType, workerName string) (host string, ok bool) {
	if d.rpcClient == nil {
		ok = d.getOneRpcClient()
		if !ok {
			rlog.V(1).Info("No disServerHost !!!")
			return
		}
	}
	if workerType != "" {
		d.workerInfo.Type = workerType
	}
	if workerName != "" {
		d.workerInfo.Name = workerName
	}
	var grt Wret
	err := d.rpcClient.Call("Drpc.GetWorkerByInfo", d.workerInfo, &grt)
	if err != nil {
		d.rpcClient = nil
		rlog.V(1).Info("Call Err:", err)
		return
	}
	host, ok = grt.Host, grt.Ok
	return
}

func (d *Discovery) Close() {
	if d.rpcClient != nil {
		d.rpcClient.Close()
	}
}

func (d *Discovery) GetWorkerListByTypeName(workerType, workerName string) (hosts []string, num int) {
	if d.rpcClient == nil {
		ok := d.getOneRpcClient()
		if !ok {
			rlog.V(1).Info("No disServerHost !!!")
			return
		}
	}
	if workerType != "" {
		d.workerInfo.Type = workerType
	}
	if workerName != "" {
		d.workerInfo.Name = workerName
	}
	var grt WlistRet
	err := d.rpcClient.Call("Drpc.GetWorkerListByTypeAndName", d.workerInfo, &grt)
	if err != nil {
		d.rpcClient = nil
		rlog.V(1).Info("Call list Err:", err)
		return
	}
	hosts, num = grt.HList, grt.Num
	return
}
