package main

import (
	"fmt"
	"log"
)

//对于 log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，
//退出程序并返回状态 1 。但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用
func test_deferfatal() {
	defer func() {
		fmt.Println("--first--")
	}()
	//log.Fatal("test for defer Fatal")
	log.Fatalln("test for defer Fatalln")
	//输出： 2019/06/12 17:24:47 test for defer Fatalln
	//log.Fatalf("test for defer Fatalf")
}

func main() {
	test_deferfatal()
}
