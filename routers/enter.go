package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.String(200, "zerotower!")
	})
	return router

}