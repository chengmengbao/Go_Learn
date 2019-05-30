package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第3个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")

	fmt.Println("slice=", string(slice)) //转换string后再打印, slice= true1234"abcgohello"

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("strconv.FormatBool(false): ", str)

	//'f'指打印格式，以小数方式，-1指小数点位数（紧缩模式），64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("strconv.FormatFloat(3.14, 'f', -1, 64): ", str)

	//整型转字符串，常用
	str = strconv.Itoa(666)
	fmt.Println("strconv.Itoa(666): ", str)
	//把字符串转换为整型
	a, _ := strconv.Atoi("567")
	fmt.Println("a=", a)

	//字符串转其它类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag=", flag)
	} else {
		fmt.Println("err=", err)
	}
}
