/**
  create by yy on 2019-08-23
*/

package main

import (
	"go_chat/app/ginServer"
	_ "go_chat/app/init"
	_ "go_chat/app/router"
)

func main() {
	ginServer.Run()
}
