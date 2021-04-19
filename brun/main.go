package main

import (
	"github.com/gzl-tommy/infra"
	"github.com/gzl-tommy/infra/base"
	_ "github.com/gzl-tommy/resk"
	"github.com/tietang/props/v3/ini"
	"github.com/tietang/props/v3/kvs"
)

func main() {
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()
}
