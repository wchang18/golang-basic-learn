package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		name string
		age  *int
		del  bool
	)

	//绑定变量
	flag.StringVar(&name, "name", "my", "姓名")
	flag.BoolVar(&del, "del", false, "删除")

	//赋值变量
	age = flag.Int("age", 0, "年龄")

	//flag注册
	flag.Parse()

	//打印
	fmt.Println("姓名:", name)
	fmt.Println("年龄:", *age)
	fmt.Println("删除:", del)
}
