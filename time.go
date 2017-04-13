package main

import "time"
import "fmt"
import "runtime"

func main() {

	go gonum()
	t()
	for {
	}
}

func t() {
	for i := 0; i < 20000; i++ {
		go thread()
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}

func gonum() {
	for {
		time.Sleep(time.Second)
		fmt.Printf("num: %d\n", runtime.NumGoroutine())
	}
}

func thread() {
	time.Sleep(10 * time.Second)
}
