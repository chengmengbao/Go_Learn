package main

import (
	_ "Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/routers"
	"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	//beego.BConfig.WebConfig.Session.SessionOn = true // 打开session 或者 在app.conf配置文件处写上sessionon = true
	beego.Run()
}
