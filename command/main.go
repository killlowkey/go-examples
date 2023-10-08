package main

import (
	"flag"
	"log"
)

var configPath string

func init() {
	// 存放配置文件
	flag.StringVar(&configPath, "c", "/app/config.yaml", "config path")
	flag.Parse()
}

// main go run main.go -c /app/test/config.yml
func main() {
	log.Println(configPath)
}
