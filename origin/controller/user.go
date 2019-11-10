package controller

import (
	_struct "go_chat/origin/struct"
	"net/http"
)

func UserLogin(res http.ResponseWriter, req *http.Request) {

	var (
		err error
	)

	// 返回 失败的json
	reply := _struct.GetError(res)

	// 解析参数
	if err = req.ParseForm(); err != nil {

		reply.Write()

		return
	}

	mobile := req.PostForm.Get("mobile")
	passWd := req.PostForm.Get("passwd")

	loginOk := false

	if mobile == "13739497421" && passWd == "123456" {
		loginOk = true
	}

	if loginOk {
		// 返回成功 json
		reply.Code = 0
		reply.Data = &_struct.LoginReply{
			Id:    1,
			Token: "test",
		}
	} else {
		reply.Msg = "user is not exists or password is incorrect"
	}

	reply.Write()

}
