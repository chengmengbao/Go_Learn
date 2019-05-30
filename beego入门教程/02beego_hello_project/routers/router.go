package routers

import (
	"Go_Learn/beego入门教程/02beego_hello_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
