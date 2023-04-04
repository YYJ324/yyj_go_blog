package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"yyj-go-blog/config"
	"yyj-go-blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		w.Done()
	}()
	w.Wait()
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	marshal, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(marshal)
	if err != nil {
		log.Println(err)
	}
}
func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	marshal, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(marshal)
	if err != nil {
		log.Println(err)
	}
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
