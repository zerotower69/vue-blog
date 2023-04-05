package config

type Upload struct {
	Size int64  `yaml:"size" json:"size"` //图片上传的大小
	Path string `yaml:"path" json:"path"` //图片上传的目录
}
