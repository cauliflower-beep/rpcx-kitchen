package main

import (
	"flag"
	"log"
	"rpcx-kitchen/chapter1/1.5transport/protocol"

	"github.com/smallnest/rpcx/server"
)

var (
	addrTcp = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith",new(protocol.Arith),"")
	_ = s.Register(new(protocol.Arith), "")
	// 服务端使用 tcp 做为网络名并且在注册中心注册了名为 serviceName/tcp@ipaddress:port 的服务
	_ = s.Serve("tcp", *addrTcp)
	log.Printf("server is listening at %v", addrTcp)
}
