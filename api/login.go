package api

import (
	"net/http"
	"yyj-go-blog/common"
	"yyj-go-blog/service"
)

func (*API) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	login, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, login)
}
