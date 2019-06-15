package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse() //用于解析命令参数，把他们的值赋给相应的变量 go run demo2.go -name=xd 执行有效
	fmt.Printf("Hello, %s!\n", name)
}
