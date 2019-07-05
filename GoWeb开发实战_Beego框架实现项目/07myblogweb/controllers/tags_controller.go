package controllers

import (
	"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/models"
	"fmt"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	//fmt.Println("tags: ", tags)
	fmt.Println(models.HandleTagsListData(tags))
	this.Data["Tags"] = models.HandleTagsListData(tags)
	this.TplName = "tags.html"
}
