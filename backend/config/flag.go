package config

import "flag"

var (
	Path = flag.String("c", "config.yaml", "配置文件路径")
)
