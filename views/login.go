package views

import (
	"net/http"
	"yyj-go-blog/common"
	"yyj-go-blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.View)
}
