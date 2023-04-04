package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}
type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Home       TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte("error"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func InitTemplate(templateDir string) HtmlTemplate {
	tp := readTemplate([]string{"index", "category", "custom", "detail", "home", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Home = tp[4]
	htmlTemplate.Login = tp[5]
	htmlTemplate.Pigeonhole = tp[6]
	htmlTemplate.Writing = tp[7]
	return htmlTemplate
}
func readTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//访问模板首页的时候，有多个模板的嵌套，解析文件的时候，需要将所有的模板都进行解析
		index := templateDir + "index.html"
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		//传入要定义的参数
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewName, index, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("解析模板出错", err)
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs
}
func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
