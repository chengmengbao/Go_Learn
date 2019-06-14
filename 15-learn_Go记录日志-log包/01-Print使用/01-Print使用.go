package main

import (
	"log"
)

func main() {
	arr := []int{2, 3, 4}
	log.Print("Print函数", arr, "\n")
	log.Println("Println函数", arr)
	log.Printf("Printf函数：%d, %d, %d\n", arr[0], arr[1], arr[2])
	//log.Printf("Printf array with item [%d,%d]\n",arr[0],arr[1])
}

//输出
//2019/06/12 17:16:25 Print函数[2 3 4]
//2019/06/12 17:16:25 Println函数 [2 3 4]
//2019/06/12 17:16:25 Printf函数：2, 3, 4
