package main

import (
	"fmt"
	//"unsafe"
)

func main() {
	fmt.Println("test")
	var a,b int
	a = 1
	c := 2
	fmt.Println(a,b,c)
	var d int = 3
	fmt.Println(d)

	var arr1 [10]int
	arr1[8] = 1
	fmt.Println(arr1)

	arr2 := [10]int{10,2}
	fmt.Println(arr2, len(arr2), cap(arr2))
	//声明数组类型时，可以不指定具体元素个数
	var arr3 []int = []int{0,1,2}
	fmt.Println(arr3, len(arr3), cap(arr3))
	//如果指定元素位置，后面未指定元素紧跟指定元素+1位置. 切片
	var arr4 []int = []int{1:3, 6:4, 5 }
	fmt.Println(arr4)
	// 使用:=声明数组变量时，必须显示声明元素个数
	// 使用var声明数组变量时，可以不声明元素个数
	var arr5 []string = []string{"fuck", "测试"}
	fmt.Println(arr5)
	arr6 := [5]int{1: 25, 345}
	fmt.Println(arr6, len(arr6), cap(arr6))
	arr7 := [...]string{"1", "2"}
	fmt.Println(arr7, len(arr7), cap(arr7))
	arr7[0] = "11"
	fmt.Println("arr7:", arr7, len(arr7), cap(arr7))
	arr8 := arr6[2:4]
	//切片后，长度指当前数组长度， cap指新切片起始位置到原数据最终的位置
	fmt.Println("arr6:", arr6, len(arr6), cap(arr6), " arr8:", arr8, len(arr8), cap(arr8))
	//切片后，新的切片指在原切片长度的基础上新切片位置至原切片的最终位置
	arr9 := arr8[1:2]
	fmt.Println("arr6:", arr6, len(arr6), cap(arr6), " arr9:", arr9, len(arr9), cap(arr9))
	
	// compare array
	var slice1 []int
	slice1 = append(slice1, arr8...)
	fmt.Println("slice1:", slice1)
	slice1 = append(slice1, 15)
	fmt.Println("slice1:", slice1)
	slice1 = append(slice1, 16,43,232)
	//copy 要求参数必须是切片不能是数组
	copy(slice1, arr2[:])
	fmt.Println("slice1:", slice1)
	
	//map dictionary (Hash Table)
	// key uniqueue
	map1 := map[string]string{"age":"12", "name":"tony"}
	fmt.Println("map1:", map1)
	map1["age"] = "34"
	fmt.Println("map1:", map1)
	delete(map1, "age")
	
	fmt.Println("map1:", map1)
	
	//function
	fmt.Println("Module/2/1: ", Module(6,4))
	fmt.Println("Module/3/1: ", Module_1(6,4,3))
	// 函数类型
	var recorder func (name string, age int, seniority int) (done bool)
	recorder = func(name string, age int, seniority int) (done bool){
		return true	
	}
	fmt.Println("function type recorder:", recorder("1",23,1))
	
	var encrypt Encipher = func(plaintext string) []byte {
		return []byte{0x16}
	}
	fmt.Println(GenEncryptionFunc(encrypt)("11"))	
	
	var myIntSlice MyIntSlice = []int{1,2,3,4}
	fmt.Println(myIntSlice.Max())
	myIntSlice.Add(15)
	fmt.Println(myIntSlice.Max())
	myIntSlice.Add1(15)
	fmt.Println(myIntSlice.Max())

	anonym := struct {
		a int
		b string
	}{0, "string"}
	fmt.Println(anonym)
	
	//struct operation
	fmt.Println(Test{Name:"tony", Address:"Beijing", Age:12})
	t := Test{}
	t.Age = 32
	t.Name = "张三"
	fmt.Println(t, t.Name)
	
	t1 := TestSub{}
	t1.Age = 12
	t1.schoolName = "北京大学"
	fmt.Println(t1)

	t2 := TestSubTwo{}
	t2.schoolName = "农业大学"
	t2.Age = 43
	fmt.Println(t2)
	
	//make() : slice,map,chan
	s := make([]int, 10, 100)
	fmt.Println(s)
	ch_s := make(chan int,10)
	ch_s <- 12
	ch_s <- 13
	fmt.Println(<-ch_s) 
	fmt.Println(<-ch_s)
	ch_s = nil
	
	// 字面量 调用内建函数new 调用内建函数make
	const pi_num = 3.1415926		
	const (
		utc1 = 6.3
		utc2, utc3 = false, "c"
		utc4, utc5
	)
	fmt.Println(utc1, utc2, utc3, utc4, utc5)
	
	const (
		ac = 2
		bc = iota
		cc 
		dc = 34
		ec
	)
	fmt.Println(ac, bc, cc, dc, ec)
	
	var switch_a = 10
	switch switch_a {
		default:
			fmt.Println("nomatch")
		case 9:
			fmt.Println("switch output is 9")
		case 10:
			fmt.Println("switch output is 10")
			fallthrough
		case 5:
			fmt.Println("switch output is 5")
	}

	

}

type TestSubTwo struct {
	TestSub
	class int
}

type TestSub struct {
	Test
	schoolName string
}

type Test struct {
	Name string `json:"name"`
	Age uint8 `json:"age"`
	Address string `json:"addr"`
}

// 带指针接收者可改变接收者自身
// 接收者是切片时不需要指定为指针
type MyIntSlice []int
func(self MyIntSlice) Max() (result int) {
	//self = append(self, 11,1,1,1,1,1)
	result = len(self)
	return
}
func(self *MyIntSlice) Add(val int) (result int) {
	*self = append(*self, val)
	result = len(*self)
	return
}
//不会改变原始变量
func(self MyIntSlice) Add1(val int) (result int) {
	self = append(self, val)
	result = len(self)
	return
}



//function
func Module(x, y int) int {
	return x % y
}

type Encipher func(plaintext string) []byte

func GenEncryptionFunc(encrypt Encipher) func(string) (ciphertext string){
	return func(plaintext string) string {
		return fmt.Sprintf("%x", encrypt(plaintext))
	}
}

func Module_1(x, y, z int) (result int) {
	result = x % y + z
	return
}



