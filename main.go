package main

import (
	"fmt"
	"go/common"
	"go/router"
	"log"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {
	fmt.Println("hello")
	serve := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// 路由功能
	router.Router()

	if err := serve.ListenAndServe(); err != nil {
		log.Print(err)
	}
}

// func index(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	var indexData IndexData
// 	indexData.Title = "go博客"
// 	indexData.Desc = "现在是入门教程"
// 	jsonStr, _ := json.Marshal(indexData)
// 	w.Write(jsonStr)
// 	fmt.Print("run in index")
// }
