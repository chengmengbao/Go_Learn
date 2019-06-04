package main

import (
	"fmt"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000/hello")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("resp.Header = ", resp.Header)
	fmt.Println("resp.Body = ", resp.Body) //是一个io流
	fmt.Println("resp.Status = ", resp.Status)
	fmt.Println("resp.StatusCode = ", resp.StatusCode)

	buf := make([]byte, 4*1024) //4k的缓存buf
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp = ", tmp)
}
