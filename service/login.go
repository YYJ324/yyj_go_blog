package service

import (
	"errors"
	"yyj-go-blog/dao"
	"yyj-go-blog/models"
	"yyj-go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "yyj")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//根据每个用户uid生成token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userinfo models.UserInfo
	userinfo.Uid = user.Uid
	userinfo.Username = user.Username
	userinfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userinfo,
	}
	return lr, nil
}
