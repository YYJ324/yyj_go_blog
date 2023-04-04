package router

import (
	"net/http"
	"yyj-go-blog/api"
	"yyj-go-blog/views"
)

func Router() {
	//       返回的类型：  1.页面 views
	//                   2.数据(json)
	//                   3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	//  http://localhost:8080/c/1
	http.HandleFunc("/c/", views.HTML.Category)
	//  http://localhost:8080/p/3.html
	http.HandleFunc("/p/", views.HTML.Detail)
	//  http://localhost:8080/writing
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/login", api.Api.Login)
	http.HandleFunc("/api/v1/post/search", api.Api.SearchPost)
	http.HandleFunc("/api/v1/post", api.Api.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.Api.GetPost)
	http.HandleFunc("/api/v1/qiniu/token", api.Api.QiniuToken)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
