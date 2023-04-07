package main

import (
	"go-blog/core"
	_ "go-blog/docs"
	"go-blog/flag"
	"go-blog/global"
	"go-blog/routers"
)

// @title vue-blog API接口文档
// @version 1.0
// @description golang开发的后台接口文档
// @host 127.0.0.1:8080
// @BasePath /api/v1/
func main() {
	//read the config file
	core.InitConf()
	//connect to mysql
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)
	global.Log = core.InitLogger() //set logger
	options := flag.Parse()
	//command line parameters
	if flag.IsWebStop(options) {
		flag.MakeMigraOptions(options)
	}
	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("server running in %s successfully!", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
