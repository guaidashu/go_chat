package router

import (
	"fmt"
	"go_chat/origin/controller"
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

		fmt.Println("register ", tplName)

		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {

			err = tpl.ExecuteTemplate(writer, tplName, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})

	}

}

func init() {
	// 设置静态文件路径
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// 1 提供静态文件支持
	http.HandleFunc("/user/login", controller.UserLogin)

	http.HandleFunc("/user/register", controller.UserRegister)

	http.HandleFunc("/user/addfriend", controller.AddFriend)

	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)

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
}
