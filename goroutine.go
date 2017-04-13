package main

// goroutine execute function or method

import (
	"fmt"
	"time"
	"os"
	"syscall"
	"os/signal"
	"sync"
	"runtime"
)

func main() {
	fmt.Println("starting......")
	
	//t1()
	//t2()
	//t3()
	//t5()
	//t6()
	t7()
}

func t7() {
	ch := make(chan int, 5)
	sign := make(chan byte, 2)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
		for{
			e, ok := <-ch
			fmt.Printf("%d (%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done.")
		sign <- 1
	}()
	<-sign
	<-sign
}

func t6() {
	pipe := make(chan int) 
	
	go func() {
		for{
			<-pipe
			fmt.Println("receive from pipeline")
		}
	}()

	go func() {
		for{
			<-pipe
			fmt.Println("receive from pipeline1")
		}
	}()



	go func() {
		for{
			//time.Sleep(time.Second)
			pipe <- 1
			fmt.Println("send to pipeline")
		}
	}()
	select{
	}

}

func t5() {
	//var ch_t chan<- int
	//var ch_t1 <-chan int
	//ch_t<- 1
	//<- ch_t1
	
}

func t4() {
	//考虑go的执行机制
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s.\n", who)
		}(name)
	}
	//runtime.Gosched()
	time.Sleep(time.Second)
}

func t3() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello, %s.\n", name)
		}()
		time.Sleep(time.Second)
	}
	runtime.Gosched()
}

func t2() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello, %s\n", name)
	}()
	time.Sleep(50000 * time.Nanosecond)
	name = "Harry"
	time.Sleep(5 * time.Second)
}

func t1() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received as signal from sigRecv1: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
		wg.Done()
	}()
	
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 2 seconds... ")
	time.Sleep(2 * time.Second)
	fmt.Printf("Stop notification... ")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	close(sigRecv2)
	fmt.Printf("done. [sigRecv1]\n")
	wg.Wait()
}

