package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-kitchen/1.4clientDemo/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &protocol.Args{
		A: 10,
		B: 20,
	}

	rep := &protocol.Reply{}
	// 同步调用
	err := xclient.Call(context.Background(), "Mul", args, rep)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, rep.C)
}
