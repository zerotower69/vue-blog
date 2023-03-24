package routers

import (
	"github.com/gin-gonic/gin"
	"go-vue-blog/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouterApp := router.Group("api/v1")
	routerGroupApp := RouterGroup{apiRouterApp}
	routerGroupApp.SettingsRouter() //settings api
	return router
}
