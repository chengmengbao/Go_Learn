package main

import "fmt"

func main() {
	// a:=10
	var p *int
	// p=&a
	p = new(int) //分配内存空间，使用完自动释放

	*p = 666
	fmt.Println("*p=", *p) // *p= 666

	q := new(int) //自动推到类型
	*q = 777
	fmt.Println("q=", q)   // q= 0xc00004c078
	fmt.Println("*q=", *q) // *q= 777
}
