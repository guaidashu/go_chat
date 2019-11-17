package main

import (
	_ "go_chat/origin/init"
	_ "go_chat/origin/router"
)

func main() {

	Run("127.0.0.1:8088")

}
