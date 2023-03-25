package models

// SysModel record system info
type SysModel struct {
	MODEL
	Version   string `gorm:"size:32" json:"version"`     //版本
	SiteTitle string `gorm:"size:32" json:"site_tile"`   //网站标题
	SiteBeiAn string `gorm:"size:32" json:"site_bei_an"` //网站备案
}
