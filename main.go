package main

import (
	_ "github.com/mohakkataria/go-docker-deploy/routers"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"fmt"
	"path/filepath"
	"path"
	"os"
	"github.com/astaxie/beego/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = false
	}
	beego.Run()
}

func init() {
	beego.BConfig.Listen.ServerTimeOut = 60
	beego.BConfig.Log.AccessLogs = true
	logs.Async()
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	path_of_executable, _ := filepath.Abs(path.Dir(os.Args[0]))
	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	viper.AddConfigPath(path_of_executable + "/conf")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}