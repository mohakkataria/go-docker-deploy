package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/mohakkataria/go-docker-deploy/routers"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
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

	pathOfExecutable, _ := filepath.Abs(path.Dir(os.Args[0]))
	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	viper.AddConfigPath(pathOfExecutable + "/conf")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
