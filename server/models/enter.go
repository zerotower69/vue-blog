package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"` //主键id
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"updated_at"'`          //更新时间
}

type PageInfo struct {
	Page  int      `json:"page" form:"page"`   //页码
	Limit int      `json:"limit" form:"limit"` //每页条数
	Key   string   `json:"key" form:"key"`     //关键字
	Sort  string   `json:"sort" form:"sort"`   //排序
	Like  []string `json:"Like" form:"like"`   //需要模糊匹配的列表
}
