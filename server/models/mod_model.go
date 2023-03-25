package models

import "go-blog/models/ctype"

// ModModel record user mod
type ModModel struct {
	MODEL
	UserID  uint        `json:"user_id"`
	User    UserModel   `gorm:"foreignKey:UserID" json:"user"` //用户
	Avatar  string      `json:"avatar"`                        //头像
	IP      string      `gorm:"size:32" json:"ip"`             //ip
	Addr    string      `gorm:"size:42" json:"addr"`
	Content string      `gorm:"size:256" json:"content"`
	Drawing ctype.Array `gorm:"type:longtext" json:"drawing"`
}
