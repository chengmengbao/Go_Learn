package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取响应内容
	buf := make([]byte, 4*1024) //4K
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("读取响应完毕")
			break
		}
		result += string(buf[:n]) //累加读取的内容
	}
	return
}

//爬取一个段子
func SpiderOneJoy(url string) (title, content string, err error) {
	resp, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	re := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re == nil {
		err = fmt.Errorf("regexp.MustCompile err re")
		return
	}
	//取题目
	tmptitle := re.FindAllStringSubmatch(resp, 1)
	for _, data := range tmptitle {
		title = data[1]
		//title = strings.Replace(title, "\r\n", "", -1)
		//title = strings.Replace(title, "\r", "", -1)
		//title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	re1 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`)
	if re1 == nil {
		err = fmt.Errorf("regexp.MustCompile err re1")
		return
	}
	//取题目
	tmpcontent := re1.FindAllStringSubmatch(resp, -1)
	for _, data := range tmpcontent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		break
	}

	return
}

func StoreJoytoFile(i int, fileTitle, fileContent []string) {
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	defer f.Close()

	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n")
		f.WriteString(fileContent[i])
		f.WriteString("\n=================================\n")

	}
}

//爬取主页
func SpiderPape(i int, page chan<- int) {

	//明确爬取的url
	//https://www.pengfu.com/xiaohua_1.html
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页网页，它的URL: %s\n", i, url)

	//开始爬取页面内容
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	//fmt.Println("result = ", result)  //打印出爬取网页的内容
	//接下来解析，取，
	// <h1 class="dp-b"><a href="https://www.pengfu.com/content_1857274_1.html"
	//解析表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//fmt.Println("joyUrls = ", joyUrls)

	//取网址
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	for _, data := range joyUrls {
		//fmt.Println("url = ", data[1])
		//开始爬取每一个笑话，每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err = ", err)
			continue
		}

		//fmt.Printf("title = #%v#", title)
		//fmt.Printf("content = #%v#", content)
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}
	//fmt.Println("fileTitle=", fileTitle)
	//fmt.Println("fileContent=", fileContent)
	StoreJoytoFile(i, fileTitle, fileContent)

	page <- i //写内容，写num
}

func DoWork(start, end int) {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)

	page := make(chan int)
	for i := start; i <= end; i++ {
		//爬取主页面的函数
		go SpiderPape(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}

}

func main() {
	var start, end int
	fmt.Printf("请输入起始页(>=1)：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(>=起始页)：")
	fmt.Scan(&end)

	//开始工作
	DoWork(start, end)
}
