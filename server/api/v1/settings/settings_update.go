package settings

import (
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/core"
	"go-blog/global"
	"go-blog/models/res"
)

func (SettingsApi) SettingsInfoUpdate(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
