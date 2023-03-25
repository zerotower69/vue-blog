package config

type SiteInfo struct {
	CreateAt    string `yaml:"create_at" json:"createAt"`
	BeiAn       string `yaml:"bei_an" json:"beiAn"`
	Title       string `yaml:"title" json:"title"`
	Version     string `yaml:"version" json:"version"`
	QQImage     string `yaml:"qq_image" json:"qqImage"`
	WechatImage string `yaml:"wechat_image" json:"wechatImage"`
	Email       string `yaml:"email" json:"email"`
	Name        string `yaml:"name" json:"name"`
	Job         string `yaml:"job" json:"job"`
	Addr        string `yaml:"addr" json:"addr"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	SloganEn    string `yaml:"slogan_en" json:"sloganEn"`
	Web         string `yaml:"web" json:"web"`
	BilibiliUrl string `yaml:"bilibili_url" json:"bilibiliUrl"`
	GiteeUrl    string `yaml:"gitee_url" json:"giteeUrl"`
	GithubUrl   string `yaml:"github_url" json:"githubUrl"`
}
