package main

import (
	_ "github.com/mohakkataria/go-docker-deploy/docs"
	_ "github.com/mohakkataria/go-docker-deploy/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
