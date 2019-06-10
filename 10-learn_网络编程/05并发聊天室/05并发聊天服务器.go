package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户 cliAddr ===》 Client
var onlineMap map[string]Client

var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体，默认，用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli
	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某个在线
	//message <- "[" + cli.Addr + "]" + cli.Name + ": login"
	message <- MakeMsg(cli, "login")
	for {

	}
}

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message //没有消息前，这里会阻塞

		//遍历变量map，给map每个成员都发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了
	//遍历变量map，给map每个成员都发送此消息
	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		//处理用户链接
		go HandleConn(conn)
	}
}
