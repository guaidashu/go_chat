package main

import (
	"go_chat/origin/controller"
	"net/http"
)

func main() {
	// 设置静态文件路径
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// 1 提供静态文件支持
	http.HandleFunc("/user/login", controller.UserLogin)

	if err := http.ListenAndServe("127.0.0.1:8087", nil); err != nil {
		panic(err)
	}
}
