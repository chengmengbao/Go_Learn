package main

import "fmt"

func test(var_map map[int]string) {
	delete(var_map, 3)
}

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m=", m)
	test(m) //map是引用传递形参
	fmt.Println("m=", m)
}
