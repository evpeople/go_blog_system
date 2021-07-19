package route

import (
	"fmt"
	"go_blog_system/route/internal/API"
	"log"
	"net/http"
	"regexp"
)

func init() {

	http.HandleFunc("/articles", API.ShowArticles)
	http.HandleFunc("/article", API.AnalyzeMethod)
	http.HandleFunc("/test/", test)
	HandleFunc("/api/", test, "[0-9]*")
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hehe")
}

func HandleFunc(url string, packHandler func(http.ResponseWriter, *http.Request, http.HandlerFunc), pattern string) {
	trueURL, trueFunc := packHandler()
	//todo 遇到的问题是无法直接运行包装函数，此函数必须要参数，而我如果有参数就不用包装器函数了

	if matchd, err := regexp.MatchString(pattern, getTrueUrl(handler)); err != nil {
		log.Default().Println(err)
	} else if !matchd {
		log.Default().Println("not matched")
	} else {
		http.DefaultServeMux.HandleFunc(url, handler)
	}

}

func packHandler(w http.ResponseWriter, r *http.Request, handlerFunc http.HandlerFunc) (trueURL string, trueFunc http.HandlerFunc) {
	return r.URL.Path, handlerFunc
}
