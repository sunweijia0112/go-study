package main

import (
	"fmt"
	"os"
)

//Unix里echo命令的一份实现，echo把它的命令行参数打印成一行
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
