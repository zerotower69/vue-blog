package config

// Config 所有的配置汇总
type Config struct {
	MySql    MySql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	QQ       QQ       `yaml:"qq"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Upload   Upload   `yaml:"upload"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	SiteInfo SiteInfo `yaml:"site-info"`
}
