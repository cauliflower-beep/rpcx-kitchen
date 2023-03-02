package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"log"
	"rpcx-kitchen/chapter1/1.5transport/protocol"
)

/*
	你可以发送 HTTP CONNECT 方法给 rpcx 服务器。 Rpcx 服务器会劫持这个连接然后将它作为TCP连接来使用。 需要注意，客户端和服务端并不使用http请求/响应模型来通信，他们仍然使用二进制协议。

	网络名称是 http， 它注册的格式是 serviceName/http@ipaddress:port。

	HTTP Connect并不被推荐。 TCP是第一选择。

	如果你想使用http 请求/响应 模型来访问服务，你应该使用网关或者http_invoke。
*/

var (
	addrHttp = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith",new(protocol.Arith),"")
	_ = s.Register(new(protocol.Arith), "")
	// 服务端使用 tcp 做为网络名并且在注册中心注册了名为 serviceName/tcp@ipaddress:port 的服务
	_ = s.Serve("http", *addrHttp)
	log.Printf("server is listening at %v", addrHttp)
}
