package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	fmt.Printf("len(ch)=%d, cap(ch)=%d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("子协程[%d]: len(ch)=%d, cap(ch)=%d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}
}
