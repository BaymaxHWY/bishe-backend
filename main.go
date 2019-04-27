package main

import (
	"bishe/backend/models"
	"github.com/astaxie/beego"
	_ "bishe/backend/routers"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = false
	models.InitDB()
	beego.Run()
}