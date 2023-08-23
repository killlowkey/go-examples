package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-example/pkg/common"
	"sync"
)

// main https://doc.rpcx.io/part1/quickstart.html
func main() {
	// #1
	d, err := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	if err != nil {
		panic(err)
	}

	// #2 封装 Arith 客户端
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3 创建方法参数
	args := &common.Args{
		A: 10,
		B: 20,
	}

	// #4 创建回复响应
	reply := &common.Reply{}

	// #5 调用 RPC 方法
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("sync: %d * %d = %d", args.A, args.B, reply.C)

	var wg sync.WaitGroup
	wg.Add(1)
	// 异步调用
	go func() {
		asyncReply := &common.Reply{}
		call, err2 := xclient.Go(context.Background(), "Mul", args, asyncReply, nil)
		if err2 != nil {
			panic(err2.Error())
		}

		// 等待调用完成
		<-call.Done
		wg.Done()
		log.Printf("async: %d * %d = %d", args.A, args.B, reply.C)
	}()

	wg.Wait()
}
