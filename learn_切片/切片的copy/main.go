package main

import "fmt"

func main() {
	srcSlice := []int{1, 2, 3, 4}
	dstSlice := []int{1, 1, 1, 1}

	copy(dstSlice, srcSlice)
	fmt.Println("dst = ", dstSlice)
}
