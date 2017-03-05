package routers

import (
	"github.com/astaxie/beego"
	"github.com/mohakkataria/go-docker-deploy/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/deploy", &controllers.DeployController{}, "post:Deploy"),
		beego.NSRouter("/deploystatus", &controllers.DeployController{}, "get:DeployStatus"),
		beego.NSRouter("/stop", &controllers.DeployController{}, "get:Stop"),
	)
	beego.AddNamespace(ns)
}
