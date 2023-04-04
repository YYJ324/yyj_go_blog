package views

import (
	"net/http"
	"yyj-go-blog/common"
	"yyj-go-blog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
