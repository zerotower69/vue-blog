package core

import (
	"fmt"
	"go-blog/config"
	"go-blog/global"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

// using the command : go get gopkg.in/yaml.v2  ---> read the yaml config file.
const ConfigFile = "settings.yaml"

func InitConf() {
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
	//fmt.Print(c.SiteInfo)
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	global.Log.Info("修改配置文件成功！")
	return nil
}
