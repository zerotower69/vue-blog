package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片的hash值，用于判断重复的图片
	Name string `gorm:"size:38" json:"name"` //图片的名称
	//NotUse        bool   //是否可以选择
	//ArticleModels []ArticleModel `gorm:"foreignKey:BannerID" json:"-"` //使用这个图片的文章列表
}
