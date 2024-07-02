package main

import "fmt"

func main() {
	var (
		sex string
		age int
	)
	fmt.Println("请输入年龄:")
	fmt.Scan(&age)
	fmt.Println("请输入性别:")
	fmt.Scan(&sex)
	fmt.Println("性别:", sex, "年龄:", age)
}
