package main

import "fmt"

func testa() {
	fmt.Println("aaa")
}
func testb() {
	fmt.Println("bbb")
	panic("this is panic")

}
func testc() {
	fmt.Println("ccc")
}
func main() {
	testa()
	testb()
	testc()
}
