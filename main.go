package main

import (
	"github.com/Godyu97/vege9/vegelog"
	"github.com/Godyu97/vegeGoWebTemplate/internal/common"
	"github.com/Godyu97/vegeGoWebTemplate/internal/config"
	"github.com/Godyu97/vegeGoWebTemplate/internal/dao"
	"github.com/Godyu97/vegeGoWebTemplate/internal/server"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

//	@title			xxx
//	@version		1.0
//	@description	golang后端平台服务项目demo
//	@contact.name	vege9
//	@contact.email	godyu97@foxmail.com
//	@license.name	Apache 2.0
//	@license.url	https://github.com/Godyu97
//	@host
//	@BasePath
func main() {
	//todo init cfg
	conf := os.Getenv("xxx_ConfFile")
	if conf == "" {
		conf = common.DefaultConfFileName
	}
	config.InitCfg("./etc", conf)
	cfg := config.GetCfg()
	//init log
	vegelog.InitLogger(path.Join("./log", "xxx_log.log"), vegelog.JudgeLogLevel(cfg.Debug))
	vegelog.LogInfo("xxx Init 🚀~")
	//init db
	//dao.InitDb(cfg.MySql, cfg.Redis)
	engine := gin.New()
	server.RegHttpRouter(engine)
	go func() {
		if err := engine.Run(cfg.ListenAddr); err != nil {
			panic(err)
		}
	}()
	common.Quit.WaitCloseWithFn(func() {
		vegelog.LogInfo("xxx Closed 🔚~")
		dao.CloseDbs()
	})
}
