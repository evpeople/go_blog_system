package API

import (
	"encoding/json"
	"fmt"
	"go_blog_system/model"
	"log"
	"net/http"
	"strconv"
)

//通过showType确定按照什么方式展示文章，也就是更改SQL语句,通过showNum，决定展示多少文章.offsetNum表明已经展示了多少文章
func ShowArticlesFromTime(w http.ResponseWriter, req *http.Request) {

	showType := req.URL.Query().Get("showType")
	show := req.URL.Query().Get("showNum")
	offset := req.URL.Query().Get("offsetNum")
	offsetNum, err := strconv.Atoi(offset)
	if err != nil {
		log.Println(err)
	}
	showNum, err := strconv.Atoi(show)
	if err != nil {
		log.Println(err)
	}
	var articleModels []model.ArticleModel

	switch showType {
	case "":
		model.DB.Offset(offsetNum).Limit(showNum).Find(&articleModels)
		data, err := json.Marshal(articleModels)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		w.Write(data)
	}
}

//show one Article
func OneArticle(w http.ResponseWriter, req *http.Request) {

}

func ShowArticlesFromCategory(w http.ResponseWriter, req *http.Request) {

}
func AddNewArticle(w http.ResponseWriter, req *http.Request) {
	var article model.ArticleModel
	var jsonData map[string]string

	if err := json.NewDecoder(req.Body).Decode(&jsonData); err != nil {
		log.Println(err)
	}
	article.Constuctor(jsonData["Author"], jsonData["Content"], jsonData["FirstPicture"])
	if result := model.DB.Create(&article); result.Error != nil {
		fmt.Println(result)

		fmt.Fprintln(w, result.Error)
	}

	fmt.Fprintf(w, "创建成功")
}
