//fibonacco 1 1 2 3 5 8 前两个数相加等于后一个数

package main

import "fmt"

//ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag=", flag)
			return
		}
	}

}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	//消费者，从channel读取内容
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		//可以停止
		quit <- true
	}() //别忘了（）
	//生产者, 产生数字，写入channel
	fibonacci(ch, quit)
}
