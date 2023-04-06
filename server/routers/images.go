package routers

import v1 "go-blog/api/v1"

func (router RouterGroup) ImagesRouter() {
	app := v1.ApiGroupApp.ImagesApi
	router.POST("images/single", app.ImageSingleUpload)
	router.POST("images/multi", app.ImageMultiUpload)
	router.GET("images/list", app.ImageList)
	router.POST("images/delete", app.ImageRemove)
	router.POST("images/update", app.ImageUpdate)
}
