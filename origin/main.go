package main

import (
	"fmt"
	"go_chat/origin/controller"
	_ "go_chat/origin/model"
	"html/template"
	"log"
	"net/http"
)

func RegisterView() {
	// 解析 template
	tpl, err := template.ParseGlob("views/**/*")

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, v := range tpl.Templates() {

		tplName := v.Name()

		fmt.Println("注册 ", tplName)

		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {

			err = tpl.ExecuteTemplate(writer, tplName, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})

	}

}

func main() {
	// 设置静态文件路径
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// 1 提供静态文件支持
	http.HandleFunc("/user/login", controller.UserLogin)

	// http.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
	//	// 解析 template
	//	tpl, err := template.ParseFiles("views/user/login.html")
	//
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	err = tpl.ExecuteTemplate(writer, "/user/login.shtml", nil)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	// })

	// file, err := os.Open("views/user/login.html")
	//
	// if err != nil {
	//	log.Fatal(err.Error())
	// }
	//
	// var buf []byte
	//
	// buf, err = ioutil.ReadAll(file)
	//
	// fmt.Println(string(buf))

	RegisterView()

	if err := http.ListenAndServe("127.0.0.1:8088", nil); err != nil {
		panic(err)
	}
}
