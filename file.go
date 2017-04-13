package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//打开文件  os.Open(string)->os.File,error;
//关闭文件 os.File.Close()->error;
// 数据缓冲区 bufio.NewReader(os.File)->*bufio.Reader,error; bufio.Reader.ReadString()->string,error;
// 字节流缓冲区 bytes.NewBuffer(io.Reader)->bytes.Buffer; bytes.Buffer.ReadForm(io.Reader)->int,error; bytes.Buffer.Bytes()->[]byte
// 快速读取字节流 io/ioutil.ReadFile(string)->[]byte,error; io/ioutil.WriteFile(string,[]byte ,mode)->error;
// gzip.NewReader(*io.Reader)->gzip.Reader,error;

// 写文件 bufio.NewWriter(io.Writer)->*bufio.Writer; bufio.Write.WriteString(string)->int,error; bufio.Write.Write([]byte)->int,error;
//    bufio.Writer.Flush()->error;

func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	f6()
}

func f1() {
	//打开文件
	file, err := os.Open("file.go")
	if err != nil {
		fmt.Printf("open failed! error msg:%s\n", err.Error())
		return
	}
	//清理关闭文件
	defer file.Close()

	//读取数据到缓冲区
	iReader := bufio.NewReader(file)

	for {
		// 从缓冲区按行读取数据
		str, err := iReader.ReadString('\n')
		if err != nil {
			return // error or EOF
		}
		fmt.Printf("The input was: %s", str)
	}

}

func f2() {
	buff, err := ioutil.ReadFile("array.go")
	fmt.Println("------------------")
	if err != nil {
		fmt.Printf("error. %s\n", err.Error())
		return
	}

	fmt.Printf("%s", buff)

	err = ioutil.WriteFile("array1.go", buff, 0x644)
	if err != nil {
		panic(err.Error())
	}

}

// gzip
func f3() {
	var r *bufio.Reader
	file, err := os.Open("t.zip")
	if err != nil {
		panic(err.Error())
	}
	fz, err := gzip.NewReader(file)
	if err != nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("read string failed!")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

// write file
func f4() {
	outputFile, outputError := os.OpenFile("output.data", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello World\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

// copy file
func f5() {
	// 局部函数定义  var funcName func(xxx,xxx)(xxx,xxx) = func(xxx,xxx)(xxx,xxx){}
	//   funcName := func(xxx,xxx)(xxx,xxx){}
	var CopyFile func(string, string) (int64, error) = func(dstName, srcName string) (written int64, err error) {
		src, err := os.Open(srcName)
		if err != nil {
			return
		}
		defer func() {
			fmt.Println("file close")
			src.Close()
		}()

		dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer dst.Close()

		return io.Copy(dst, src)
	}

	len, err := CopyFile("array2.go", "array.go")
	fmt.Println("here file is closed!")
	if err != nil {
		fmt.Println("copy failed!")
	}
	fmt.Printf("copy file size: %d\n", len)
}

func f6() {
	who := "Alice"
	//os.Args 获取命令参数长度 默认是1
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:2], "  ")

	}
	fmt.Println("Good Morning", who)
}
