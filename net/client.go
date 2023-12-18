package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	// 从标准输入中读取数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// 将数据发送到服务器
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
