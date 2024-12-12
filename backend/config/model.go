package config

var Config = &config{}

type config struct {
	Port        int    `yaml:"port"`
	LoggerModel string `yaml:"logger"`
	Database    struct {
		Connection  string `yaml:"connection"`    //数据库连接信息
		Driver      string `yaml:"driver"`        //数据库驱动
		MaxOpenCons int    `yaml:"MAX_OPEN_CONS"` //数据库最大打开连接数
		MaxIdleCons int    `yaml:"MAX_IDLE_CONS"` //数据库最大空闲连接数
		MaxLifeTime int    `yaml:"MAX_LIFE_TIME"` //数据库连接最大生命周期
	} `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}
