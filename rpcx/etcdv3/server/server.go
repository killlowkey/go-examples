package main

import (
	"flag"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
	"net/http"
	"rpcx-example/pkg/common"
	"time"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func init() {
	flag.Parse()
}

// https://github.com/rpcxio/rpcx-examples/blob/master/registry/etcdv3/server/server.go
func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":9981", nil))
	}()

	s := server.NewServer()
	addRegistryPlugin(s)

	log.Println("Register Arith rpc service...")
	err := s.RegisterName("Arith", new(common.Arith), "")
	if err != nil {
		panic(err)
	}

	log.Println("Start Rpc service successfully")
	err = s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,      // rpc 服务地址
		EtcdServers:    []string{*etcdAddr}, // etcd 注册中心地址
		BasePath:       *basePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start() // 连接到 etcd
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
