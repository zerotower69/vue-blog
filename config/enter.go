package config

// Config define the config struct
type Config struct {
	MySql  MySql  `yaml:"mySql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
