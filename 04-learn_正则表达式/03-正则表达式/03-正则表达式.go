package main

import (
	"fmt"
	"regexp"
)

func main() {
	//`` 反引号代表原生字符串
	buf := `
		<div>哈哈1</div>
		<div>哈哈2</div>
		<div>哈哈3</div>
		<div>哈哈4</div>
	`

	//1.解析规则，它会解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`<div>>.*</div>`)
	if reg1 == nil {
		fmt.Println("解析失败")
		return
	}

	//2.根据规则提取关键信息
	// res1 := reg1.FindAllString(buf, -1)   //res1= [3.14 1.23 8.99 6.66]
	res1 := reg1.FindAllStringSubmatch(buf, -1) //res1= [[3.14] [1.23] [8.99] [6.66]]
	fmt.Println("res1=", res1)
}
