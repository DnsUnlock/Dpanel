package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadeConfig() error {
	// 读取配置文件
	data, err := os.ReadFile(*Path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, Config)
	if err != nil {
		return err
	}
	return nil
}
