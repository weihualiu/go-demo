package main

import "errors"
import "fmt"
import "time"

func main() {
	go test_g()
	fmt.Println(test())	
	time.Sleep(5*time.Second)
	fmt.Println("end")	
}

func test() error {
	return errors.New("fuck your mother!!!")
}

func test_g() {
	time.Sleep(time.Second)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered panic: %s\n", r)
		}
	}()
	test_g_p()
}

func test_g_p() {
	panic(errors.New("goroutine error"))
}


