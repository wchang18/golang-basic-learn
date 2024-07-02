package main

import (
	"fmt"
	"os"
)

func main() {

	//创建文件
	os.Create("file.txt")

	//打开文件
	file, err := os.Open("file.txt")
	fmt.Println(err)
	state, _ := file.Stat()
	fmt.Println(state.Size())

	//打开一个文件
	//os.O_RDWR:读写模式
	//os.O_APPEND:追加模式
	//0666:读写权限
	file2, err := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND, 0666)
	fmt.Println(file2.WriteString("aaa"))
	fmt.Println(file2.WriteString("bbb"))

	// file的write方法
	file3, err := os.Create("file2.txt")
	fmt.Println(file3.Write([]byte("ccc")))
	//读取文件
	fmt.Println(os.ReadFile("file2.txt"))
	//写入文件
	fmt.Println(os.WriteFile("file2.txt", []byte("ddd"), 0666))
}
