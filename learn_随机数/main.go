package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需一次
	// 若果种子参数一样，每次运行程序产生的随机数都一样
	// rand.Seed(666)
	fmt.Println("time.Now().UnixNano():", time.Now().UnixNano())

	rand.Seed(time.Now().UnixNano()) // 以当前系统时间作为种子参数

	for i := 0; i < 5; i++ {
		fmt.Println("rand = ", rand.Int())     // 产生随机数，且很大
		fmt.Println("rand = ", rand.Intn(100)) // 限制在100内的数
	}

}
