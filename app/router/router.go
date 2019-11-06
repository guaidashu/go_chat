/**
  create by yy on 2019-07-02
*/

package router

import (
	"fmt"
	"go_chat/app/controller"
	"go_chat/app/ginServer"
	"net/http"
)

func init() {
	fmt.Println("开始初始化router")

	ginServer.LoadHTMLGlob("app/views/index/*")

	test := ginServer.Group("/test")
	{
		test.GET("/", controller.Test)
	}
	user := ginServer.Group("/user")
	{
		user.GET("/login", controller.UserLogin)
	}
	ginServer.GET("/", controller.Index)
	ginServer.GET("/index", controller.Index)
	ginServer.GET("/init_table", controller.InitTables)

	http.FileServer()

	fmt.Println("router初始化成功")
}
