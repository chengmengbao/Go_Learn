package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	oldCap := cap(s)
	fmt.Println("s= ", s)
	fmt.Println("oldCap= ", oldCap)

	s = append(s, 1)
	fmt.Println("s= ", s)
	fmt.Println("1 newCap= ", cap(s))

	s = append(s, 2)
	fmt.Println("s= ", s)
	fmt.Println("2 newCap= ", cap(s))

	s = append(s, 3)
	fmt.Println("s= ", s)
	fmt.Println("3 newCap= ", cap(s))

	// s = append(s, 4)
	// fmt.Println("s= ", s)
	// fmt.Println("4 newCap= ", cap(s))
}
