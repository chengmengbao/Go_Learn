package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("文件读取错误：", err)
		return
	}
	fmt.Println("1 文件的内容：", data)
	fmt.Println("2 文件的内容：", string(data))
}
