package main

import "fmt"
import "io"
import "net"

func main() {
	fmt.Println(".......")

	var (
		host          = "182.207.129.67"
		port          = "21"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)

	conn, err := net.Dial("tcp", remote)
	io.WriteString(conn, msg)
	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))

	}
	conn.Close()
}
