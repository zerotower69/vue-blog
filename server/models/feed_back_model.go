package models

type FeedBackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"` //email
	Content      string `json:"content"`              //反馈的内容
	ApplyContent string `json:"apply_content"`        //回复的内容
	IsApply      bool   `json:"is_apply"`             //是否回复
}
