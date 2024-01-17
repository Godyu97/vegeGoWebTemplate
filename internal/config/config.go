package config

import (
	"github.com/Godyu97/vege9/vegeConfig"
)

type Cfg struct {
	ProjectName string
	ListenAddr  string
	MySql       DSN
	JwtKey      JwtKeyConf
	Redis       RedisConf
	Debug       bool
}

type DSN struct {
	Username string
	Password string
	Address  string
	Database string
}

type RedisConf struct {
	Address  string
	Password string
	Database int
}

type JwtKeyConf struct {
	ActJwtKey string
	RetJwtKey string
}

var cfgObj *Cfg

// 返回值类型，无法被外部更改
func GetCfg() Cfg {
	if cfgObj != nil {
		return *cfgObj
	} else {
		panic("mBYKPPhT need InitCfg")
	}
}

func InitCfg(dir string, file string) {
	cfgObj = new(Cfg)
	vegeConfig.InitYamlCfg(dir, file, cfgObj)
}
