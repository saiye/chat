package service

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetCfg(dir string, fileName string, fileType string) (*viper.Viper, error) {
	cfg, err := NewConfig(dir, fileName, fileType)
	if err != nil {
		fmt.Errorf("读取配置文件异常.. %v", err)
		return nil, err
	}
	return cfg, err
}

func NewConfig(dir string, fileName string, fileType string) (*viper.Viper, error) {
	if dir == "" {
		dir = "./"
	}
	//1.读取配置文件
	config := viper.New()
	config.AddConfigPath(dir)      // 文件所在目录
	config.SetConfigName(fileName) // 文件名
	config.SetConfigType(fileType) // 文件类型
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}
	return config, nil
}
