package util

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var cfg *goconfig.ConfigFile

/**
*	初始化配置文件信息的读取
 */
func InitConf() error {
	var err error
	cfg, err = goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("LoadConfigFile confi.ini fail, %v", err)
		return err
	}
	return nil
}

/**
*获取配置文件中sec章节中key的值
 */
func GetValue(sec string, key string) (string, error) {
	value, err := cfg.GetValue(sec, key)
	return value, err
}
