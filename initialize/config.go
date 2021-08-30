package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-demo/global"
	"go.uber.org/zap"
)

func InitConfig() {
	//配置目录拼接
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-dev.yaml", configFilePrefix)

	v := viper.New()

	//文件路径设置
	v.SetConfigFile(configFileName)
	//读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//反序列化将配置文件变成数据结构
	if err := v.Unmarshal(global.LocalConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置信息: &v", global.LocalConfig)

}
