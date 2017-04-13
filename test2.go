package main

import "fmt"
import "time"

func main() {
	fmt.Println("test")
	ch := make(chan int, 100)
	flag := make(chan bool)
	f1(ch, flag)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond)
			ch <- i
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Millisecond)
			ch <- i
		}
	}()

	<-flag
}

//flag标记通道
func f1(ch11 <-chan int, flag chan<- bool) {
	go func() {
		var e int
		ok := true
		for {
			select {
			case e, ok = <-ch11:
				fmt.Printf("channel %d\n", e)
			case ok = <-func() chan bool {
				timeout := make(chan bool, 1)
				go func() {
					time.Sleep(2 * time.Millisecond)
					timeout <- false
				}()
				return timeout
			}():
				fmt.Println("Timeout")
				flag <- true
				break
			}
			if !ok {
				break
			}
		}
	}()
}
