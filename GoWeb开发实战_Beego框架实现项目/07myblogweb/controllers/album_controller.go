package controllers

import (
	//"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/models"
	//"github.com/opentracing/opentracing-go/log"
	"Go_Learn/GoWeb开发实战_Beego框架实现项目/07myblogweb/models"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		//log.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}
