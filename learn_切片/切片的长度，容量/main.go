package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[:]
	fmt.Println("s1=", s1)
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))

	s2 := a[3:6:7]
	fmt.Println("s2=", s2)
	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))

	s3 := a[:6]
	fmt.Println("s3=", s3)
	fmt.Printf("len=%d, cap=%d\n", len(s3), cap(s3))

	s4 := a[3:]
	fmt.Println("s4=", s4)
	fmt.Printf("len=%d, cap=%d\n", len(s4), cap(s4))
}
