package main

import "fmt"
import "strings"

func main() {
	//strings.Contains,源字符串是否包含子串，返回bool值
	fmt.Println("strings.Contains: ", strings.Contains("helloworld", "hello")) //strings.Contains:  true
	fmt.Println("strings.Contains: ", strings.Contains("helloworld", "rorl"))  //strings.Contains:  false

	//strings.Join,
	s := []string{"hello", "world"}
	buf := strings.Join(s, ", ")
	fmt.Println("buf=", buf) //buf= hello, world

	//strings.Index  返回所在位置的索引，若没有则返回-1
	fmt.Println("strings.Index=", strings.Index("helloworld", "go")) //strings.Index= -1

	//strings.Repeat
	fmt.Println("strings.Repeat=", strings.Repeat("go", 4)) //strings.Repeat= gogogogo

	//strings.Split 以指定的分隔符拆分
	buf = "abc,edf,gjh"
	s1 := strings.Split(buf, ",")
	fmt.Println("s1=", s1) //s1= [abc edf gjh]

	//strings.Fields与Split类似，不过它只能以空格作为分隔,返回一个切片
	s2 := strings.Fields("    hello world    ")
	fmt.Println("s2=", s2) //s2= [hello world]

	//strings.Trim去掉两头的字符
	buf = strings.Trim(",,,,helloworld,,,,,,", ",") //去掉两头的,
	fmt.Printf("buf=#%s#\n", buf)                   //buf=#helloworld#

}
