package main

import (
	"fmt"
	_ "go_blog_system/model"
	_ "go_blog_system/route"
	"net/http"
)

func main() {
	if err := http.ListenAndServe("127.0.0.1:3456", nil); err != nil {
		fmt.Println(err)
	}
}
