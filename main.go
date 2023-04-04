package main

import (
	"yyj-go-blog/common"
	"yyj-go-blog/server"
)

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//程序入口
	server.App.Start("127.0.0.1", "1111")
}
