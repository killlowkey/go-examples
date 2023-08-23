package main

import (
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
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

func main() {
	// 创建服务并添加 etcd 注册插件
	s := server.NewServer()
	addRegistryPlugin(s)

	// 注册 rpc 服务
	err := s.RegisterName("Arith", new(common.Arith), "")
	if err != nil {
		panic(err)
	}

	// 启动服务
	go func() {
		log.Fatal(s.Serve("tcp", *addr))
	}()

	time.Sleep(time.Minute)

	// 解除注册服务
	err = s.UnregisterAll()
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
