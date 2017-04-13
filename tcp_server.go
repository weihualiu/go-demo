package main

import (
	//"errors"
	"fmt"
	"net"
)

func main() {
	fmt.Println("network......")
	//建立监听
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	//循环接收客户端来的请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		//开启一个goroutine处理新连接
		go doServerStuff(conn)
	}

}

func doServerStuff(conn net.Conn) {
	//循环接收数据
	for {
		buf := make([]byte, 1024)
		//将网络连接数据读取到字节数组变量中
		_, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("connection interrupt")
				return
			}
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("receive data: %v\n", string(buf))

	}
}
