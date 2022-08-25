package router

import (
	"go/api"
	"go/views"
	"net/http"
)

func Router() {
	//1.页面 views 2.数据 api (json) 3.静态资源

	http.HandleFunc("/", views.HTML.Index)

	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login/", views.HTML.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/login", api.API.Login)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
