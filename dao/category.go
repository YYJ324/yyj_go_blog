package dao

import (
	"log"
	"yyj-go-blog/models"
)

func GetAllCategory() ([]models.Category, error) {
	query, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory查询出错", err)
		return nil, err
	}
	var categorys []models.Category
	for query.Next() {
		var category models.Category
		err := query.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory取值出错", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
