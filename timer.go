package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------")
	t := time.NewTimer(2 * time.Second)
	now := time.Now()
	fmt.Printf("Now time:%v.\n", now)

	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)

	<-time.NewTimer(3 * time.Second).C

}
