package main

import (
	"context"
	"flag"
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-example/pkg/common"
	"time"
)

var (
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	d, _ := etcd_client.NewEtcdDiscovery(*basePath, "Arith", []string{*etcdAddr}, false, nil)
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := &common.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &common.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(5 * time.Second)
	}
}
