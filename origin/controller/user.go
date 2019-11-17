package controller

import (
	"fmt"
	"go_chat/app/libs"
	"go_chat/origin/model"
	"go_chat/origin/service"
	_struct "go_chat/origin/struct"
	"math/rand"
	"net/http"
)

func UserLogin(res http.ResponseWriter, req *http.Request) {

	var (
		err error
	)

	// 解析参数
	if err = req.ParseForm(); err != nil {

		_struct.WriteError(res, libs.NewReportError(err).Error())

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
		_struct.WriteSuccess(res, &_struct.LoginReply{
			Id:    1,
			Token: "test",
		})
	} else {
		_struct.WriteError(res, "user is not exists or password is incorrect")
	}

}

func UserRegister(res http.ResponseWriter, req *http.Request) {

	var (
		err         error
		userService *service.UserService
		user        *model.UserModel
	)

	userService = &service.UserService{}

	// 解析参数
	if err = req.ParseForm(); err != nil {

		_struct.WriteError(res, libs.NewReportError(err).Error())

		return
	}

	mobile := req.PostForm.Get("mobile")

	plainPwd := req.PostForm.Get("passwd")

	nickName := fmt.Sprintf("user%06d", rand.Int31())

	avatar := ""
	sex := model.SEX_UNKONW

	if user, err = userService.Register(mobile, plainPwd, nickName, avatar, sex); err != nil {
		_struct.WriteError(res, libs.NewReportError(err).Error())
		return
	}

	_struct.WriteSuccess(res, user)

}
