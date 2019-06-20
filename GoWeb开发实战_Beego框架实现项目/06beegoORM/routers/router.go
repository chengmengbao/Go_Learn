package routers

import (
	"beegoORM/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testcreatetable", &controllers.ModelController{}, "get:CreateTable")
	beego.Router("/testmodel", &controllers.ModelController{}, "get:Get")
	beego.Router("/testquery", &controllers.ModelController{}, "get:Query;post:Post")
}
