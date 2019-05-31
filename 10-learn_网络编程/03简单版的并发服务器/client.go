package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//发起请求连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//主函数结束，关闭连接
	defer conn.Close()

	//处理标准输入
	go func() {
		buf := make([]byte, 2048) //2k的字节切片
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.Stdin.Read err = ", err)
				return
			}

			conn.Write(buf[:n])
		}
	}()

	//处理接收服务器发来的响应信息
	str := make([]byte, 2048)
	for {
		n, err := conn.Read(str)
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		fmt.Println(string(str[:n]))
	}

}
