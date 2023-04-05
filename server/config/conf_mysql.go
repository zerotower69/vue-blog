package config

import "strconv"

// MySql define mysql config struct
type MySql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级，debug就是输出全部的sql
}

// Init 初始化值
func (t *MySql) init() {
	t.Host = "127.0.0.1"
	t.Port = 3306
	t.Config = ""
	t.DB = "blog"
	t.User = "root"
	t.Password = "123456"
	t.LogLevel = "debug"
}

// Dsn 获取MySQL链接地址
func (t *MySql) Dsn() string {
	return t.User + ":" + t.Password + "@tcp(" + t.Host + ":" + strconv.Itoa(t.Port) + ")/" + t.DB + "?" + t.Config
}
