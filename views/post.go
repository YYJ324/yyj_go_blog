package views

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"yyj-go-blog/common"
	"yyj-go-blog/service"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	// 1.找到返回的页面
	detail := common.Template.Detail
	// 获取路径参数
	path := r.URL.Path
	pid_str := strings.TrimPrefix(path, "/p/")
	//7.html
	pid_str = strings.TrimSuffix(pid_str, ".html")
	pId, err := strconv.Atoi(pid_str)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	//根据请求路径获取到了post_id,根据这个id拿数据
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
