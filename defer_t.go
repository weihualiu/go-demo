package main

import "fmt"

func main() {
	//f()
	//fmt.Println("Returned normally from f.")
	record()
}

func f() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered in f", r)
	//	}
	//}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Println in g", i)
	g(i + 1)
}


func begin(funcName string) string {
	fmt.Printf("Enter function %s.\n", funcName)
	return funcName
}

func end(funcName string) string {
	fmt.Printf("Exit function %s.\n", funcName)
	return funcName
}

func record() {
	// 在record函数执行时，defer中的函数参数会被初始化为静态值
	defer end(begin("record"))
	fmt.Println("In function record.")
}

