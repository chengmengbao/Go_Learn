package main

import "fmt"

type long int

func Add(a, b int) int {
	return a + b
}
func (a long) Add01(b long) long {
	return a + b
}

func main() {
	res := Add(1, 1)
	fmt.Println("res = ", res)

	var a long = 2
	r := a.Add01(3)
	fmt.Println("r = ", r)
}



