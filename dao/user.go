package dao

import (
	"log"
	"yyj-go-blog/models"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select username from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}
func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow(
		"select * from blog_user where username=? and passwd=? limit 1",
		userName,
		passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.Username, &user.Passwd, &user.Avatar, &user.UpdateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
