package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	View   Viewer
	System SystemConfig
}

var Cfg *tomlConfig

func init() {
	//程序启动的时候，就会执行init方法
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "yyj-go-blog"
	Cfg.System.Version = 1.0
	cur_dir, _ := os.Getwd()
	Cfg.System.CurrentDir = cur_dir
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}
