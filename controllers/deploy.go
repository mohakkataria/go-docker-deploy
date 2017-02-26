package controllers

import (
	"github.com/mohakkataria/go-docker-deploy/payloads"
	_ "encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about object
type DeployController struct {
	beego.Controller
}

func (this *DeployController) Deploy() {
	var requestObj payloads.Deploy

	if len(this.Ctx.Request.Form) == 0 {
		this.Data["json"] = map[string]interface{}{"message": "Empty request", "status": "failed"}
		this.ServeJSON()
		return
	}

	if err := this.ParseForm(&requestObj); err != nil {
		this.Data["json"] = map[string]interface{}{"message": "Invalid request format", "status": "failed"}
		this.ServeJSON()
		return
	}

	logs.Info("[Request][Deploy] : ", requestObj)
}

func (this *DeployController) DeployStatus() {
	var dockerName string
	if err := this.Ctx.Input.Bind(&dockerName, "name"); err != nil {
		logs.Info("[Request][DeployStatus] : ", err)
		this.Data["json"] = map[string]interface{}{"message": "No Docker Name specified", "status": "failed"}
		this.ServeJSON()
		return
	} else {
		logs.Info("[Request][DeployStatus] : ", dockerName)
	}


}

func (this *DeployController) Stop() {

	var dockerName string
	if err := this.Ctx.Input.Bind(&dockerName, "name"); err != nil {
		logs.Info("[Request][Stop] : ", err)
		this.Data["json"] = map[string]interface{}{"message": "No Docker Name specified", "status": "failed"}
		this.ServeJSON()
		return
	} else {
		logs.Info("[Request][Stop] : ", dockerName)
	}


}