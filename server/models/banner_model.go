package models

import (
	"go-blog/global"
	"go-blog/models/ctype"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片的hash值，用于判断重复的图片
	Name      string          `gorm:"size:38" json:"name"`         //图片的名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片的类型，本地还是七牛
	//NotUse        bool   //是否可以选择
	//ArticleModels []ArticleModel `gorm:"foreignKey:BannerID" json:"-"` //使用这个图片的文章列表
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//本地图片，删除还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
