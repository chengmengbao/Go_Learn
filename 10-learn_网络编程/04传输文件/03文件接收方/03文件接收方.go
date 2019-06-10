package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf) //读取对方发送的文件名
	if err2 != nil {
		fmt.Println("conn.Read err2 = ", err2)
		return
	}

	filename := string(buf[:n])
	fmt.Println("filename = ", filename)

	conn.Write([]byte("ok")) //回复“Ok”

	RecvFile(filename, conn)
}

func RecvFile(filename string, conn net.Conn) {
	f, err := os.Create(filename) //新建文件
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	//接收多少，写多少，一点不差
	for {
		n, err1 := conn.Read(buf) //接收对方发过来的文件内容
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read = ", err1)
			}
			return
		}

		//if n == 0 {
		//	fmt.Println("n == 0, 文件读取完毕")
		//	break
		//}

		f.Write(buf[:n]) //往文件写入内容
	}

}
