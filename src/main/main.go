package main

import (
	"github.com/spf13/viper"
	"golang-starter-project/src/infra"
	// 触发全局init函数
	_ "golang-starter-project/src"
)

// 全局Starter启动器，对Starter进行配置并启动项目
func main() {
	// 获取全局配置文件
	conf := getConfig()
	app := infra.NewBootApplication(conf)
	app.Run()
	select {}
}

func getConfig() *viper.Viper {
	conf := viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath("./src/main/")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
