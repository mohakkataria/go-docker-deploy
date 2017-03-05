package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/mitchellh/mapstructure"
	"github.com/mohakkataria/go-docker-deploy/models"
	"github.com/spf13/viper"
)

// DeployController - Controller for deploy apis
type DeployController struct {
	beego.Controller
}

// Deploy - to deploy a docker image using the json input
func (dc *DeployController) Deploy() {
	var requestObj models.DockerContainer

	reqObj := dc.Ctx.Input.RequestBody
	if err := json.Unmarshal(reqObj, &requestObj); err != nil {
		dc.Data["json"] = map[string]interface{}{"message": "Invalid input data : " + err.Error(), "status": "failed"}
		dc.ServeJSON()
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

	dc.Data["json"] = returnDeployStatuses
	dc.ServeJSON()
}

// DeployStatus -- Check the deploy status of docker container
func (dc *DeployController) DeployStatus() {
	var dockerName string
	if err := dc.Ctx.Input.Bind(&dockerName, "name"); err != nil {
		logs.Info("[Request][DeployStatus] : ", err)
		dc.Data["json"] = map[string]interface{}{"message": "No Docker Name specified", "status": "failed"}
		dc.ServeJSON()
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

	dc.Data["json"] = returnDeployStatuses
	dc.ServeJSON()

}

// Stop -- stop a docker
func (dc *DeployController) Stop() {

	var dockerName string
	if err := dc.Ctx.Input.Bind(&dockerName, "name"); err != nil {
		logs.Info("[Request][Stop] : ", err)
		dc.Data["json"] = map[string]interface{}{"message": "No Docker Name specified", "status": "failed"}
		dc.ServeJSON()
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

	dc.Data["json"] = returnStopStatuses
	dc.ServeJSON()

}
