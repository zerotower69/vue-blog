package images

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/res"
)

func (ImagesApi) ImageRemove(c *gin.Context) {
	var cr models.RemoveRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)

	res.OkWihMessage(fmt.Sprintf("共删除 %d 张图片", count), c)
}
