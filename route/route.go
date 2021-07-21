package route

import (
	"github.com/gorilla/mux"
	"go_blog_system/route/internal/API"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	rp := r.Methods("POST").Subrouter()
	//articles 展示所有的文章 ，通过 query参数辨别应该展示的部分
	//article  id 展示对应id的文章
	//article post 上传文章
	r.HandleFunc("/api/v1/articles/", API.ShowArticlesFromTime)
	r.HandleFunc("/api/v1/articles/{category}/{id:[0-9]+}", API.OneArticle)
	r.HandleFunc("/api/v1/articles/{category}/", API.ShowArticlesFromCategory)

	//所有post方法的路由
	rp.HandleFunc("/api/v1/articles/", API.AddNewArticle)
	http.Handle("/", r)
}
