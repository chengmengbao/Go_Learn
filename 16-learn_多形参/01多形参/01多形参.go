package main

import (
	"fmt"
)

func TestMultiParam() {
	valueArray := []string{"1", "2", "3", "4", "5"}
	result := valueArray[0:3]
	fmt.Println(result)

	multiParam(result...) // 这里就是我们平时需要用到的
}

func multiParam(args ...string) {
	fmt.Println(args)
}

func main() {
	TestMultiParam()
}
