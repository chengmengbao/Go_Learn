package main

import "fmt"

func main() {

	// var a [5]int = [5]int{1, 2, 3, 4, 5}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("a[%d] = %d\n", i, a[i])
	// }

	//指定某个元素初始化
	// d := [5]int{2: 10, 4: 20}
	// fmt.Println("d=", d)
	a := [5]int{1, 2, 3, 4, 5}
	var d [5]int
	d = a
	fmt.Println("d=", d)
	fmt.Printf("&a = %p, a = %v\n", &a, a)
	fmt.Printf("&d = %p, d = %v\n", &d, d)
}
