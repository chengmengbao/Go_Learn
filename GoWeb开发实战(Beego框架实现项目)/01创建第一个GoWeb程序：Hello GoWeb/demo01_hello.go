package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
它需要2个参数，一个是http.ResponseWriter，另一个是http.Request
往http.ResponseWriter写入什么内容，浏览器的网页源码就是什么内容。
http.Request里面是封装了，浏览器发过来的请求（包含路径、浏览器类型等等）。
*/
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello 世界！"))
	fmt.Println("------------------")
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path：", r.URL.Path)
	fmt.Println("scheme：", r.URL.Scheme)
	fmt.Println(r.Form["url_ long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	// fmt.Fprintf(w, "Hello GoWeb!") //这个写入到w的是输出到客户端的
}

/*
1.运行该程序
2.打开浏览器，输入：http://localhost:8080
*/
func main() {
	/**
	  第一个参数：pattern string,
	  第二个参数：handler func(ResponseWriter, *Request)
	*/
	http.HandleFunc("/", sayHello) //设置访问的路由

	/**
	    第一个参数addr：监听地址
		第二个参数handler：通常为空，意味着服务端调用http.DefaultServerMux进行处理，
		而服务端编写的业务逻辑处理程序http.Handle()或http.HandleFunc()默认注入http.DefaultServeMux中
	*/
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
