package controllers

import (
	_ "github.com/mohakkataria/go-docker-deploy/payloads"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/spf13/viper"
	"github.com/mohakkataria/go-docker-deploy/models"
	"github.com/mitchellh/mapstructure"
)

// Operations about object
type DeployController struct {
	beego.Controller
}

func (this *DeployController) Deploy() {
	var requestObj models.DockerContainer

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
	}

	logs.Info("[Request][DeployStatus] : ", dockerName)

	dockerContainer := models.DockerContainer{
		Name:dockerName,
	}

	var host models.Host
	results := make(chan map[string]string, 10)
	defer close(results)

	hostsFromConfig := viper.Get("hosts").([]interface{})
	for _, hostFromConfig := range hostsFromConfig {
		mapstructure.Decode(hostFromConfig, &host)
		go func(host models.Host) {
			results <- dockerContainer.GetDeploymentStatus(host)
		}(host)

	}

	returnDeployStatuses := []map[string]string{}

	for i := 0; i < len(hostsFromConfig); i++ {
		select {
		case res := <-results:
			returnDeployStatuses = append(returnDeployStatuses, res)
		}
	}

	this.Data["json"] = returnDeployStatuses
	this.ServeJSON()

}

func (this *DeployController) Stop() {

	var dockerName string
	if err := this.Ctx.Input.Bind(&dockerName, "name"); err != nil {
		logs.Info("[Request][Stop] : ", err)
		this.Data["json"] = map[string]interface{}{"message": "No Docker Name specified", "status": "failed"}
		this.ServeJSON()
		return
	}

	logs.Info("[Request][Stop] : ", dockerName)

	dockerContainer := models.DockerContainer{
		Name:dockerName,
	}

	var host models.Host
	results := make(chan map[string]string, 10)
	defer close(results)

	hostsFromConfig := viper.Get("hosts").([]interface{})
	for _, hostFromConfig := range hostsFromConfig {
		mapstructure.Decode(hostFromConfig, &host)
		go func(host models.Host) {
			results <- dockerContainer.Stop(host)
		}(host)

	}

	returnStopStatuses := []map[string]string{}

	for i := 0; i < len(hostsFromConfig); i++ {
		select {
		case res := <-results:
			returnStopStatuses = append(returnStopStatuses, res)
		}
	}

	this.Data["json"] = returnStopStatuses
	this.ServeJSON()


}