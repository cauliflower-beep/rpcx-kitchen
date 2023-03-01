package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-kitchen/1.3serverDemo/protocol"
)

var (
	addr2 = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &protocol.Args{
		A: 10,
		B: 20,
	}

	rep := &protocol.Reply{}
	call, err := xclient.Go(context.Background(), "Mul", args, rep, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	repCall := <-call.Done
	if repCall.Error != nil {
		log.Fatalf("fai;ed to call: %v", repCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, rep.C)
	}
}