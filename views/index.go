package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"yyj-go-blog/common"
	"yyj-go-blog/service"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err2 := r.ParseForm(); err2 != nil {
		log.Println("表单获取出错:", err2)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	index.WriteData(w, hr)
}
