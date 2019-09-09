package base

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init()  {
	//定义日子格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	formatter.ForceFormatting = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:"green",
		WarnLevelStyle:"yellow",
		TimestampStyle: "37",
	})
	logrus.SetFormatter(formatter)
	//日志级别
	level := os.Getenv("log.debug")
	if level == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	//控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false
	//日志文件和滚动配置
	logrus.Info("测试")
	logrus.Debug("cesi")
}

