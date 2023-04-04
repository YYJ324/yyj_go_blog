package views

import (
	"net/http"
	"yyj-go-blog/common"
	"yyj-go-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
