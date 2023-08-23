package main

import (
	"github.com/smallnest/rpcx/server"
	"log"
	"rpcx-example/pkg/common"
)

// main https://doc.rpcx.io/part1/quickstart.html
func main() {
	s := server.NewServer()
	_ = s.RegisterName("Arith", new(common.Arith), "")
	log.Println("Start RpcX service success")
	log.Fatal(s.Serve("tcp", ":8972"))
}
