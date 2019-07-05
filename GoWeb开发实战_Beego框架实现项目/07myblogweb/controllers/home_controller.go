package controllers

import (
	"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/models"
	"fmt"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

func (this *HomeController) Get() {

	tag := this.GetString("tag")
	fmt.Println("tag:", tag)

	page, _ := this.GetInt("page")
	fmt.Println("page=", page)

	var artList []models.Article

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}

		artList, _ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

	this.TplName = "home.html"
}
