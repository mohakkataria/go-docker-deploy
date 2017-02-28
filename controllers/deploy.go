package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/mitchellh/mapstructure"
	"github.com/mohakkataria/go-docker-deploy/models"
	_ "github.com/mohakkataria/go-docker-deploy/payloads"
	"github.com/spf13/viper"
)

// Operations about object
type DeployController struct {
	beego.Controller
}

func (this *DeployController) Deploy() {
	var requestObj models.DockerContainer

	reqObj := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(reqObj, &requestObj); err != nil {
		this.Data["json"] = map[string]interface{}{"message": "Invalid input data : " + err.Error(), "status": "failed"}
		this.ServeJSON()
		return
	}

	logs.Info("[Request][Deploy] : ", requestObj)

	var host models.Host
	results := make(chan map[string]string, 10)
	defer close(results)

	hostsFromConfig := viper.Get("hosts").([]interface{})
	for _, hostFromConfig := range hostsFromConfig {
		mapstructure.Decode(hostFromConfig, &host)
		go func(host models.Host) {
			results <- requestObj.Deploy(host)
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
		Name: dockerName,
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
		Name: dockerName,
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
