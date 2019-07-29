package myconfig

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// MyConf 导出配置结构体
type MyConf struct {
	Concurrent int    `yaml:"concurrent"` // 并发量，根据具体情况写，我在一台机器上写 100W，结果机子卡死了
	IPPort     string `yaml:"ipPort"`
	TryTimes   int    `yaml:"tryTimes"`
}

// GetConf 获取配置文件指针
func (c *MyConf) GetConf() *MyConf {

	yamlFile, err := ioutil.ReadFile("myconfig.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
