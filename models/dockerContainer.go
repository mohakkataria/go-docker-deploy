package models

import (
	"github.com/mohakkataria/go-docker-deploy/managers/sshManager"
	"strings"
)

type DockerContainer struct {
	Name string `json:"name"`
	Image     string   `json:"image"`
	Ports     string   `json:"ports"`
	EnvironmentVars     string   `json:"environment_vars"`
	Command     string   `json:"command"`
	Volumes 	string `json:"volumes"`
}

func (this *DockerContainer) Deploy(host Host) (map[string]string) {
	command := "docker run -d -it "

	if (this.Name != "") {
		command += " --name " + this.Name
	}

	for _,value := range strings.Split(this.Volumes, ";") {
		if value != "" {
			command = command + " -v " + value
		}
	}

	for _,value := range strings.Split(this.Ports, ";") {
		if value != "" {
			command = command + " -p " + value
		}
	}

	if (this.Image != "") {
		command += " " + this.Image
	} else {
		return map[string]string{
			"host" : host.Ip,
			"output" : "Missing Image",
		}
	}

	if (this.Command != "") {
		command += "  " + this.Command
	}
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