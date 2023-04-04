package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"yyj-go-blog/common"
	"yyj-go-blog/dao"
	"yyj-go-blog/models"
	"yyj-go-blog/service"
	"yyj-go-blog/utils"
)

func (*API) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pid_str := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pid_str)
	if err != nil {
		common.Error(w, errors.New("不识别此路径"))
		return
	}
	post, err := dao.GetPostByPid(pId)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}
func (*API) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户id，判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//POST  save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// update
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}

func (*API) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	sreachRes := service.SearchPost(condition)
	common.Success(w, sreachRes)
}
