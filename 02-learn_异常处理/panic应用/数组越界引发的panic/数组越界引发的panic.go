package main

import "fmt"

func testa() {
	fmt.Println("aaa")
}
func testb(x int) {
	// fmt.Println("bbb")
	// panic("this is panic")
	var a [5]int
	a[x] = 111 //数组越界
}
func testc() {
	fmt.Println("ccc")
}
func main() {
	testa()
	testb(5)
	testc()
}
