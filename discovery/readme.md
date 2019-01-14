### etcd 用于服务发现的用法说明

### cleint 端使用说明
     在对应的客户端中引用
     "wacaispider/discovery"
     示例：
     在主函数 main() 执行下 main的代码
     discovery.NewRegister("server/stock","login","GTJA","192.168.3.101:6372",[]string{"http://192.168.3.108:2379","http://192.168.3.109:2379","http://192.168.3.99:2379"})
     
    其中：
    server/stock 指服务的大类别 对应于etcd的key前缀
    login        服务的大类别
    GTJA         指服务的具体类别
    192.168.3.101:6372  为本服务的开启服务host
    []string{"192.168.3.108:3379","192.168.3.109:3379","192.168.3.99:3379"} 为disServer 的配置参数数组

### server服务端的使用
    server端可以根据 server.go中的示例代码进行集体修改
    WatchWorkers
    server端Watch 
    得到的key: /server/stock/login/GTJA/192.168.3.101:6374
    数据内容: {"Prekey":"server/stock","Host":"172.16.6.30:7878","Type":"login","Name":"GJZQ","Sid":"54fd82de508c588735eb4aba10564b52"}
	
### 编译环境中的说明
    
