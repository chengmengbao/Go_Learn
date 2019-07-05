package controllers

import (
	"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/models"
	"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/utils"
	"fmt"
	"strconv"
)

type ShowArticleController struct {
	//beego.Controller
	BaseController
}

func (this *ShowArticleController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)

	this.Data["Title"] = art.Title
	//this.Data["Content"] = art.Content
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName = "show_article.html"
}
