package settings

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfo(c *gin.Context) {
	var cr SettingsUri

	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
		break
	//case "mysql":
	//	res.OkWithData(global.Config.MySql, c)
	//	break
	case "email":
		res.OkWithData(global.Config.Email, c)
		break
	//case "logger":
	//	res.OkWithData(global.Config.Logger, c)
	//	break
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
		break
	case "qq":
		res.OkWithData(global.Config.QQ, c)
		break
	case "system":
		res.OkWithData(global.Config.System, c)
		break
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
		break
	default:
		res.FailWithMessage(`没有对应的配置信息`+":"+cr.Name, c)
	}
}
