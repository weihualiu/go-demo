package main

import "fmt"

//import "time"

func main() {
	unbufChan := make(chan int)
	var num int = 2
	fmt.Printf("num is %d.\n", num)
	go func() {
		fmt.Println("Sleep a second...")
		//time.Sleep(time.Millisecond)
		//num := <-unbufChan
		fmt.Printf("Received a integer %d.\n", num)
		unbufChan <- num
	}()
	//time.Sleep(time.Second)
	num = 1
	fmt.Printf("Send integer %d...\n", num)
	//unbufChan <- num
	<-unbufChan
	fmt.Println("Done.")
}
