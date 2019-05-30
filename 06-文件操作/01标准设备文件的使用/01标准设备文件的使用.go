package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stdout.WriteString("hhhh\n")

	var a int
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
