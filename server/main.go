package main

import (
	"go-vue-blog/core"
	"go-vue-blog/global"
	"go-vue-blog/routers"
)

func main() {
	//read the config file
	core.InitConf()
	//connect to mysql
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)
	global.Log = core.InitLogger() //set logger
	//global.Log.Warn("warning hahahah")
	//global.Log.Infof("zerotower good job!")
	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("server running in %s successfully!", addr)
	router.Run(addr)
}
