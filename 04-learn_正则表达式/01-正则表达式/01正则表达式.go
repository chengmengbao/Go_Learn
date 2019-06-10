package main

import (
	"fmt"
	"regexp"
)

func main() {

	// buf := "abc azc a7c aac 888 a9c tac"
	buf := "abcazca7caac888a9ctac"
	//1.解析规则，它会解析正则表达式，如果成功返回解释器
	// reg1 := regexp.MustCompile(`a.c`)
	// reg1 := regexp.MustCompile(`a[0-9]c`)
	reg1 := regexp.MustCompile(`a[\d]c`)
	if reg1 == nil {
		fmt.Println("解析失败")
		return
	}

	//2.根据规则提取关键信息
	res1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("res1=", res1) //res1= [[a7c] [a9c]]
}
