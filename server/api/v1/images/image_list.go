package images

import (
	"github.com/gin-gonic/gin"
	"go-blog/models"
	"go-blog/models/res"
	"go-blog/service/common"
)

// ImageList 图片列表查询
func (ImagesApi) ImageList(c *gin.Context) {
	//var imageList []models.BannerModel
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr) // 将query解析并传给cr
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//TODO:根据是否生产环境和日志等级决定是否需要打印
	list, total, err := common.GetOffsetList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})

	////计算总数
	//total := global.DB.Select("id").Find(&imageList).RowsAffected
	////计算偏移量
	//offset := (cr.Page - 1) * cr.Limit
	//if offset < 0 {
	//	offset = 0
	//}
	////分页查询 ==> select * from banner_models limit cr.Limit offset offset
	//global.DB.Limit(cr.Limit).Offset(offset).Find(&imageList)
	res.OkWithPage(list, total, c)
}
