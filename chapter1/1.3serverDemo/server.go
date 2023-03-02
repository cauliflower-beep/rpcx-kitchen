package main

import (
	"flag"
	"log"
	"rpcx-kitchen/chapter1/1.3serverDemo/protocol"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.RegisterName("Arith",new(protocol.Arith),"")
	_ = s.Register(new(protocol.Arith), "")
	_ = s.Serve("tcp", *addr)
	log.Printf("server is listening at %v", addr)
}
