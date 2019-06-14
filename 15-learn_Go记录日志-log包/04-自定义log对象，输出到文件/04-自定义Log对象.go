package main

import (
	"log"
	"os"
)

//func init() {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//}

func main() {
	//配合func init()使用
	//log.Println("飞雪无情的博客:", "http://www.flysnow.org")
	//log.Printf("飞雪无情的微信公众号：%s\n", "flysnow_org")

	fileName := "test.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	//该函数一共有三个参数：
	//（1）输出位置out，是一个io.Writer对象，该对象可以是一个文件也可以是实现了该接口的对象。
	// 通常我们可以用这个来指定日志输出到哪个文件。
	//（2）prefix 我们在前面已经看到，就是在日志内容前面的东西。我们可以将其置为 "[Info]" 、
	// "[Warning]"等来帮助区分日志级别。
	//（3） flags 是一个选项，显示日志开头的东西，可选的值有：
	//Ldate         = 1 << iota     // 形如 2009/01/23 的日期
	//Ltime                         // 形如 01:23:23   的时间
	//Lmicroseconds                 // 形如 01:23:23.123123   的时间
	//Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
	//Lshortfile                    // 文件名和行号: d.go:23
	//LstdFlags     = Ldate | Ltime // 日期和时间
	//Loger := log.New(logFile, "[Info]", log.LstdFlags|log.Lshortfile)
	Loger := log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.LUTC) //如果设置了LUTC的话，就会把输出的日期时间转为0时区的日期时间显示
	Loger.Println("A Info message here")
	Loger.SetPrefix("[Debug]")
	Loger.Println("A Debug Message here")
}
