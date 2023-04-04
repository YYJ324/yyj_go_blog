package server

import (
	"log"
	"net/http"
	"yyj-go-blog/router"
)

var App = &YyjServer{}

type YyjServer struct {
}

func (*YyjServer) Start(ip, port string) {
	//web 服务，http协议：ip，port
	server := http.Server{
		Addr: ip + ":" + port,
	}
	// 路由请求
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}