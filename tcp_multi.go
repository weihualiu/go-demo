package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}

}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: `"+hostAndPort+"`")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	fmt.Println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {

	commForm := conn.RemoteAddr().String()
	fmt.Println("Connection from: ", commForm)
	sayHello(conn)
	for {
		var ibuf = make([]byte, maxRead+1)
		//从网络缓存中读取指定长度数据
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			//读取到指定长度数据进行处理
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:
			//数据读取完毕时继续进行循环
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	fmt.Println("Closed connection: ", commForm)
	checkError(err, "Close: ")
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 'f', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+string(wrote)+"bytes.")

}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		fmt.Print("<", length, ":")
		for i := 0; i < length; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		fmt.Println(">")
	}
}

func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error())

	}
}
