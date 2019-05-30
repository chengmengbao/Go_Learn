package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "3.14 567 adsf 1.23 7. 8.99 fsdfsdf 6.66 "

	//1.解析规则，它会解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`\d+\.\d+`)
	if reg1 == nil {
		fmt.Println("解析失败")
		return
	}

	//2.根据规则提取关键信息
	// res1 := reg1.FindAllString(buf, -1)   //res1= [3.14 1.23 8.99 6.66]
	res1 := reg1.FindAllStringSubmatch(buf, -1) //res1= [[3.14] [1.23] [8.99] [6.66]]
	fmt.Println("res1=", res1)
}
