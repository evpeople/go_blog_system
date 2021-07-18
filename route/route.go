package route

import (
	"fmt"
	"go_blog_system/route/internal/API"
	"net/http"
)

func init() {

	http.HandleFunc("/articles", API.ShowArticles)
	http.HandleFunc("/article", API.AnalyzeMethod)
	http.HandleFunc("/test", test)
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hehe")
}
