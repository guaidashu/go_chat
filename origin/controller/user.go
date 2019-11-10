package controller

import (
	"encoding/json"
	"fmt"
	"go_chat/app/libs"
	_struct "go_chat/origin/struct"
	"net/http"
)

func UserLogin(res http.ResponseWriter, req *http.Request) {

	var (
		result []byte
		err    error
	)

	// 设置返回数据格式
	res.Header().Set("Content-Type", "application/json")

	// 返回 失败的json
	reply := _struct.GetError()
	reply.Data = &_struct.LoginReply{
		Id:    1,
		Token: "test",
	}

	// 解析参数
	if err = req.ParseForm(); err != nil {

		res.WriteHeader(http.StatusInternalServerError)

		if result, err = json.Marshal(reply); err != nil {
			fmt.Println(libs.NewReportError(err).Error())
		}

		_, err = res.Write(result)

		if err != nil {
			fmt.Println(libs.NewReportError(err).Error())
		}

		return
	}

	mobile := req.PostForm.Get("mobile")
	passWd := req.PostForm.Get("passwd")

	loginOk := false

	if mobile == "13739497421" && passWd == "123456" {
		loginOk = true
	}

	if loginOk {

		// 设置成功状态码
		res.WriteHeader(http.StatusOK)

		// 返回成功 json
		reply.Code = 1
	}

	if result, err = json.Marshal(reply); err != nil {
		fmt.Println(libs.NewReportError(err).Error())
	}

	// 设置输出
	if _, err = res.Write(result); err != nil {
		fmt.Println(libs.NewReportError(err).Error())
	}

}
