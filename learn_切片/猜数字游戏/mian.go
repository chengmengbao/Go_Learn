package main

import (
	"fmt"
	"math/rand"
	"time"
)

func creatRandNum(p *int) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000) //一定是4位数
		if num >= 1000 {
			break
		}
	}
	*p = num
}

func getNum(randSlice []int, randNum int) {
	randSlice[0] = randNum / 1000
	randSlice[1] = randNum % 1000 / 100
	randSlice[2] = randNum % 100 / 10
	randSlice[3] = randNum % 10
}

func onGame(randSlice []int) {
	var keyNum int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Println("请输入4位数字：")
			fmt.Scan(&keyNum)
			if keyNum > 999 && keyNum < 10000 {
				break
			}
			fmt.Println("请输入正确的4位数字！")
		}
		fmt.Println("keyNum=", keyNum)
		getNum(keySlice, keyNum)
		fmt.Println("keySlice=", keySlice)
		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一点\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("全部才对，猜对数字为", keyNum)
			break
		}
	}
}

func main() {
	var randNum int

	// 产生一个4位的随机数
	creatRandNum(&randNum)
	fmt.Println("randNum=", randNum)

	randSlice := make([]int, 4)
	//获取各位的数字
	getNum(randSlice, randNum)
	fmt.Println("randSlice=", randSlice)

	onGame(randSlice)
}
