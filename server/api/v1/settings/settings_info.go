package settings

import (
	"github.com/gin-gonic/gin"
	"go-vue-blog/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OK(c)
}
