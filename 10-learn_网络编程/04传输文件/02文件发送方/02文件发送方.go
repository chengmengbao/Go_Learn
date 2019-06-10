package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}

	//注意关闭文件
	defer f.Close()

	read_buf := make([]byte, 4*1024)

	for {
		n, err1 := f.Read(read_buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件读取完毕")
			} else {
				fmt.Println("f.Read err1 = ", err1)
			}
			return
		}
		_, err2 := conn.Write([]byte(read_buf[:n]))
		if err2 != nil {
			fmt.Println("conn.Write err2 = ", err2)
			return
		}
	}

}

func main() {
	//获取用户输入的文件路径
	fmt.Println("请输入文件的绝对路径：")
	var path string
	fmt.Scan(&path)

	//获取文件属性
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	//客户端发起请求连接
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}

	//记得函数结束关闭conn
	defer conn.Close()

	//发送文件名至服务器,   字符串string转字节切片[]byte是[]byte(info.Name())
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err2 = ", err2)
		return
	}

	//接收服务器的响应
	recv_buf := make([]byte, 1024)
	n, err3 := conn.Read(recv_buf)
	if err3 != nil {
		fmt.Println("conn.Read err3 = ", err3)
		return
	}
	//如果接收的内容是ok，说明对方准备好了
	if "ok" == string(recv_buf[:n]) {
		//开始传输文件
		sendFile(path, conn)
	} else {
		fmt.Println("接收内容不是ok")
	}

}
