package core

import (
	"fmt"
	"go-vue-blog/config"
	"go-vue-blog/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// using the command : go get gopkg.in/yaml.v2  ---> read the yaml config file.

func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml config file error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init Unmarshal: %v", err)
	}
	log.Println("config yaml file loaded Init successfully!")
	global.Config = c
	//fmt.Println(c)
}
