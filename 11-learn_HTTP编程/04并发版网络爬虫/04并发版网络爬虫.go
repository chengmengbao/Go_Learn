package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)

	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页Body内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}

		res += string(buf[:n])
	}
	return
}

func SpiderPage(i int, page chan<- int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬第%d页网页%s\n", i, url)

	//爬（将所有的网站的内容全部爬下来）
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	//把内容写入到文件
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}
	f.WriteString(result) //写内容
	f.Close()             //关闭文件
	page <- i
}

func DoWork(s, e int) {
	fmt.Printf("正在爬取%d到%d的页面\n", s, e)
	//明确目标，要知道你准备在哪个范围或者网站去搜索
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0 //下一页+50
	page := make(chan int)
	for i := s; i <= e; i++ {
		go SpiderPage(i, page) //新建协程来爬
	}
	for i := s; i <= e; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var s, e int
	fmt.Println("请输入起始页（ >= 1）:")
	fmt.Scan(&s)
	fmt.Println("请输入终页（>=起始页）：")
	fmt.Scan(&e)

	DoWork(s, e)
}
