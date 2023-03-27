package config

// MySql define mysql config struct
type MySql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_Level"` //日志等级，debug就是输出全部的sql
}
