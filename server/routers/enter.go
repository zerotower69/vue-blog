package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go-blog/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouterApp := router.Group("api/v1")
	routerGroupApp := RouterGroup{apiRouterApp}
	routerGroupApp.SettingsRouter() //settings api
	routerGroupApp.ImagesRouter()   //images api
	return router
}
