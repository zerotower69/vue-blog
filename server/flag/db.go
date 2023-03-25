package flag

import (
	"go-blog/global"
	"go-blog/models"
)

func MakeMigraOptions(option Option) {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.AdvertModel{},
			&models.ArticleModel{},
			&models.BannerModel{},
			&models.CommentModel{},
			&models.FeedBackModel{},
			&models.MenuBannerModel{},
			&models.MenuModel{},
			&models.MessageModel{},
			&models.ModModel{},
			&models.SysModel{},
			&models.TagModel{},
			&models.UserCollectModel{},
			&models.UserModel{},
			&models.LogDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
