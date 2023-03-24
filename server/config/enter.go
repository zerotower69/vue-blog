package config

// Config define the config struct
type Config struct {
	MySql  MySql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
