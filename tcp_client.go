package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client......")
	//客户端连接服务器
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字:")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("请输入要扯淡的内容:")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedClient + " 说: " + trimmedInput))
		if err != nil {
			fmt.Println("connection send message failed!")
			return
		}
	}

}
