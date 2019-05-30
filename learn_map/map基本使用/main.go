package main

import "fmt"

func main() {
	//定义一个变量，类型为map [int]string
	var m1 map[int]string
	fmt.Println("m1=", m1) //m1= map[]

	//对于map只有len,没有cap
	fmt.Println("len=", len(m1)) //len= 0

	//可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2=", m2)       //m2= map[]
	fmt.Println("len=", len(m2)) //len= 0
	//可以通过make创建,可以指定长度，只是指定了容量，但是里面却是一个数据也没有
	m3 := make(map[int]string, 2)
	fmt.Println("m3=", m3)       //m2= map[]
	fmt.Println("len=", len(m3)) //len= 0

	m3[1] = "a"
	m3[2] = "b"
	m3[3] = "c"
	fmt.Println("m3=", m3)       //m3= map[3:c 1:a 2:b] 无序的
	fmt.Println("len=", len(m3)) //len= 3

	//初始化
	//键值是唯一的
	m4 := map[int]string{1: "Mike", 2: "go", 3: "python"}
	fmt.Println("m4=", m4) //m4= map[1:Mike 2:go 3:python]
}
