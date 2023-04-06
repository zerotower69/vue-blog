package common

import (
	"go-blog/global"
	"go-blog/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// GetOffsetList 获取分页的数据
func GetOffsetList[T any](model T, option Option) (list []T, total int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MySQLLog})
	}

	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间往前排序
	}

	total = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, total, err
}
