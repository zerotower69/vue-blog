package models

import (
	"go-blog/models/ctype"
)

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32;comment:'文章标题'" json:"title"`            //文章标题
	Abstract      string         `gorm:"comment:'文章概要'" json:"abstract;"`                //文章简介
	Content       string         `gorm:"comment:'文章内容'"  json:"content"`                 //文章内容
	LookCount     int            `gorm:"comment:'浏览量'"  json:"look_count"`               //浏览量
	CommentCount  int            `gorm:"comment:'评论数'" json:"comment_count"`             //评论量
	LikeCount     int            `json:"like_count"`                                     //点赞量
	CollectsCount int            `json:"collects_count"`                                 //收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"` //文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`                  //文章的评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`                     //文章作者
	UserID        uint           `json:"user_id"`                                        //用户id
	Category      string         `gorm:"size:20" json:"category"`                        //文章分类
	Source        string         `json:"source"`                                         //文章来源
	Link          string         `json:"link"`                                           //原文链接
	Word          int            `json:"word"`                                           //文章字数
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`                   //文章封面
	BannerID      uint           `json:"banner_id"`                                      //文章封面ID
	Nickname      string         `gorm:"size:42" json:"nickname"`                        //发布文章的用户昵称
	BannerPath    string         `json:"banner_path"`                                    //文章的封面
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`                //文章标签
}
