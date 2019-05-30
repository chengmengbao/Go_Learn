package main

import "fmt"

func main() {
	//创建一个双向通道
	ch := make(chan int)
	//生产者，生产数字，写入channel
	//新开一个协程
	go produce(ch) //channel传参，引用

	//消费者，从channel读取内容，打印
	comsuem(ch)

}

//此通道只能写，不能读
func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + i
	}
	close(ch)
}

//此channel只能读，不能写
func comsuem(ch <-chan int) {
	for num := range ch {
		fmt.Println("num=", num)
	}
}
