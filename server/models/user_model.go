package models

import (
	"go-blog/models/ctype"
)

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36;comment:'用户昵称'" json:"nick_name"`                                              //昵称
	Username       string           `gorm:"size:36;comment:'用户名'" json:"username"`                                                //用户名
	Password       string           `gorm:"size:128;comment:'密码'" json:"password"`                                                //密码
	Avatar         string           `gorm:"size:256;comment:'用户的头像'" json:"avatar_id"`                                            //头像id
	Email          string           `gorm:"size:128;comment:'邮箱'" json:"email"`                                                   //邮箱
	Tel            string           `gorm:"size:16;comment:'手机号'" json:"tel"`                                                     //手机号
	Addr           string           `gorm:"size:64;comment:'地址'" json:"addr"`                                                     //地址
	Token          string           `gorm:"size:64;comment:'当前用户token'" json:"token"`                                             //其他平台的唯一id
	IP             string           `gorm:"size:20;comment:'用户IP,为注册IP'" json:"ip"`                                               //IP地址
	Role           ctype.Role       `gorm:"size:4;default:1;comment:'用户角色'" json:"role"`                                          //权限 1 管理员 2 普通用户 3 游客
	SignStatus     ctype.SignStatus `gorm:"type:smallint(6);comment:'用户的注册来源'" json:"sign_status"`                                //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                           //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;joinReference:ArticleID" json:"-"` //收藏的文章列表
}
