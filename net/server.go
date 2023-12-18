package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听在本地端口 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("TCP server is running on port 8080")

	for {
		// 接受连接
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建一个新的 reader 用于接收客户端发送的数据
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Message received: ", string(message))
	}
}
