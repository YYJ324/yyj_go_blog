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

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	//拿页面
	categoryTemplate := common.Template.Category
	//  http://localhost:8080/c/1
	path := r.URL.Path
	category_id := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(category_id)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	if err2 := r.ParseForm(); err2 != nil {
		log.Println("表单获取出错:", err2)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
