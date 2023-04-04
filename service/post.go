package service

import (
	"html/template"
	"log"
	"yyj-go-blog/config"
	"yyj-go-blog/dao"
	"yyj-go-blog/models"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByPid(pid)
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.View,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil
}
func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.View.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
	}
	wr.Categorys = category
	return
}
func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetPostByCondition(condition)
	var searchResp []models.SearchResp
	for _, post := range posts {
		searchResp = append(searchResp, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResp
}
