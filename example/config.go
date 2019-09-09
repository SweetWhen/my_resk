package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "myresk"
	"myresk/infra"
)

func main() {
	//获取程序运行文件所在的路径
	fpath := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(fpath)
	app := infra.New(conf)
	app.Start()

}
