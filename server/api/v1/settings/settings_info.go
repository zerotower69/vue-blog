package settings

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
