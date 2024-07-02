package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Hostname())                           //获取主机名
	fmt.Println(fmt.Sprintf("%#v", os.Environ()))        //获取环境变量
	fmt.Println(os.Setenv("my_os", "gary"))              //设置环境变量
	fmt.Println(os.Getenv("my_os"))                      //获取环境变量
	fmt.Println(os.Getuid())                             //获取用户id
	fmt.Println(os.Getgid())                             //获取用户组id
	fmt.Println(os.Getpid())                             //获取进程id
	fmt.Println(os.Mkdir("mktest", os.ModeDir))          //创建目录
	fmt.Println(os.MkdirAll("mktest2/test", os.ModeDir)) //创建目录
	fmt.Println(os.Remove("mktest"))                     //删除目录
	fmt.Println(os.RemoveAll("mktest2"))                 //删除目录
	fmt.Println(os.Rename("file1.txt", "file2.txt"))     //重命名文件
}
