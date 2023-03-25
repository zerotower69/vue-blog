package main

import (
	"go-blog/core"
	"go-blog/flag"
	"go-blog/global"
	"go-blog/routers"
)

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
