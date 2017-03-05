package models

import (
	"github.com/mohakkataria/go-docker-deploy/managers/sshManager"
	"strings"
)

//DockerContainer -- struct for the DockerContainer properties
type DockerContainer struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	Ports           string `json:"ports"`
	EnvironmentVars string `json:"environment_vars"`
	Command         string `json:"command"`
	Volumes         string `json:"volumes"`
}

//Deploy - Method to deploy a container into a host
func (dc *DockerContainer) Deploy(host Host) map[string]string {
	command := "docker run -d -it "

	if dc.Name != "" {
		command += " --name " + dc.Name
	}

	for _, value := range strings.Split(dc.Volumes, ";") {
		if value != "" {
			command = command + " -v " + value
		}
	}

	for _, value := range strings.Split(dc.Ports, ";") {
		if value != "" {
			command = command + " -p " + value
		}
	}

	if dc.Image != "" {
		command += " " + dc.Image
	} else {
		return map[string]string{
			"host":   host.IP,
			"output": "Missing Image",
		}
	}

	if dc.Command != "" {
		command += "  " + dc.Command
	}
	output, err := sshManager.ExecuteCmd(command, host.IP, host.Port, host.PrivateKeyFilePath, host.User)
	if err != nil {
		return map[string]string{
			"host":   host.IP,
			"output": err.Error(),
		}
	}

	return map[string]string{
		"host":   host.IP,
		"output": output,
	}

}

//GetDeploymentStatus - Method to get deployment status of a container into a host
func (dc *DockerContainer) GetDeploymentStatus(host Host) map[string]string {
	command := "docker ps --filter \"name=" + dc.Name + "\" --format \"{{.Status}}\""
	output, err := sshManager.ExecuteCmd(command, host.IP, host.Port, host.PrivateKeyFilePath, host.User)
	if err != nil {
		return map[string]string{
			"host":   host.IP,
			"output": err.Error(),
		}
	}

	splitOutput := strings.Split(output, "\n")

	if len(splitOutput) > 1 {
		output = splitOutput[0]
	} else {
		output = "Container Not Running"
	}

	return map[string]string{
		"host":   host.IP,
		"output": output,
	}

}

//Stop - Method to stop a container on a host
func (dc *DockerContainer) Stop(host Host) map[string]string {
	command := "docker rm -f " + dc.Name + "; echo $?"
	output, err := sshManager.ExecuteCmd(command, host.IP, host.Port, host.PrivateKeyFilePath, host.User)
	if err != nil {
		return map[string]string{
			"host":   host.IP,
			"output": err.Error(),
		}
	}

	if strings.Contains(output, dc.Name) {
		output = "Container stopped"
	} else {
		output = "No container by that name was running"
	}
	return map[string]string{
		"host":   host.IP,
		"output": output,
	}

}
