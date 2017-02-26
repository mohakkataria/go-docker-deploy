package models

import (
	"github.com/mohakkataria/go-docker-deploy/managers/sshManager"
	"strings"
	_ "github.com/astaxie/beego/logs"
)

type DockerContainer struct {
	Name string `form:"name"`
	Image     string   `form:"image"`
	Ports     string   `form:"ports"`
	EnvironmentVars     string   `form:"environment_vars"`
	Command     string   `form:"command"`
}

func (this *DockerContainer) Deploy(host Host) (map[string]string) {
	command := "docker run -d "
	output, err := sshManager.ExecuteCmd(command, host.Ip, host.Port, host.PrivateKeyFilePath, host.User)
	if err != nil {
		return map[string]string{
			"host" : host.Ip,
			"output" : err.Error(),
		}
	} else {
		return map[string]string{
			"host" : host.Ip,
			"output" : output,
		}
	}
}


func (this *DockerContainer) GetDeploymentStatus(host Host) (map[string]string) {
	command := "docker ps --filter \"name="+this.Name+"\" --format \"{{.Status}}\""
	output, err := sshManager.ExecuteCmd(command, host.Ip, host.Port, host.PrivateKeyFilePath, host.User)
	if err != nil {
		return map[string]string{
			"host" : host.Ip,
			"output" : err.Error(),
		}
	} else {
		splitOutput := strings.Split(output, "\n")

		if len(splitOutput) > 1 {
			output = splitOutput[0]
		} else {
			output = "Container Not Running"
		}

		return map[string]string{
			"host" : host.Ip,
			"output" : output,
		}
	}
}

func (this *DockerContainer) Stop(host Host) (map[string]string) {
	command := "docker rm -f "+this.Name+"; echo $?"
	output, err := sshManager.ExecuteCmd(command, host.Ip, host.Port, host.PrivateKeyFilePath, host.User)
	if err != nil {
		return map[string]string{
			"host" : host.Ip,
			"output" : err.Error(),
		}
	} else {
		if strings.Contains(output, this.Name) {
			output = "Container stopped"
		} else {
			output = "No container by that name was running"
		}
		return map[string]string{
			"host" : host.Ip,
			"output" : output,
		}
	}
}