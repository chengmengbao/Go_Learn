package main

import "fmt"

func main() {
	a := [5]int{1, -2, 2, -4, 5}

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] < a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	fmt.Println("a=", a)
}
