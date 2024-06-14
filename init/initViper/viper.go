package initViper

import (
	"github.com/spf13/viper"
)

type Config struct {
	MYSQL struct {
		Host     string
		Port     int
		User     string
		Password string
		Db       string
	}
	REDIS struct {
		Host     string
		Port     int
		Password string
	}
}

func InitViper() *Config {
	// 初始化viper
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败: " + err.Error())
	}
	var c Config
	if err = viper.Unmarshal(&c); err != nil {
		panic("映射配置文件失败: " + err.Error())
	}
	return &c
}
