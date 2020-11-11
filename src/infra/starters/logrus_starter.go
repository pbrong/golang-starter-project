package starters

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"golang-starter-project/src/infra"
	"os"
)

// 实现BaseStarter
type LogStarter struct {
	infra.BaseStarter
}

func (l *LogStarter) Setup(conf *viper.Viper) {
	// 定义日志格式
	formatter := &prefixed.TextFormatter{}
	//开启完整时间戳输出和时间戳格式
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	// 控制台高亮显示
	formatter.ForceFormatting = true
	formatter.ForceColors = true
	formatter.DisableColors = false
	//设置高亮显示的色彩样式
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	logrus.SetFormatter(formatter)
	//日志级别，通过环境变量来设置
	// 后期可以变更到配置中来设置
	level := os.Getenv("log.debug")
	if level == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Info("测试log starter [1]")
	logrus.Debug("测试log starter [2]")
}
