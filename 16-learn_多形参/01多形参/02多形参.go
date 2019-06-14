package main

import (
	"fmt"
	"reflect"
)

func main() {

	hello("variable parameter", 1, 3, 4, 45, 56, 3)
	hello("second", []int{22, 33, 44}...) //切片作为形参时，结尾需要加...
}

func hello(b string, args ...int) {
	fmt.Println(b)
	fmt.Println(args)
	fmt.Println(reflect.TypeOf(args))
}

//输出
/*
variable parameter
[1 3 4 45 56 3]
[]int
second
[22 33 44]
[]int
*/
