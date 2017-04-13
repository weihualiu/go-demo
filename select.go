package main

import "fmt"

//import "time"

func main() {
	//f1()
	f2()
}

func f1() {
	var ch1 chan int
	var ch2 chan string

	ch1 = make(chan int, 10)
	ch1 <- 1

	select {
	default:
		fmt.Println("no case selected")
	case <-ch1:
		//<-ch1
		fmt.Println("1th case is selected.")
	case ch2 <- "1":
		fmt.Println("2th case is selected.")

	}

}

func f2() {
	var i, j, k int
	i = 1
	j = 2
	k = 3
	j, i, k = i, k, j
	fmt.Printf("i=%d, j=%d, k=%d\n", i, j, k)
}
