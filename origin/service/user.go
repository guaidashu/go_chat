package service

import (
	"errors"
	"fmt"
	"go_chat/app/libs"
	"go_chat/origin/model"
	"go_chat/origin/utils"
	"math/rand"
)

type UserService struct {
}

// 注册函数
func (s *UserService) Register(mobile, plainPwd, nickName, avatar, sex string) (user *model.UserModel, err error) {

	var tmpUser *model.UserModel

	userModel := new(model.UserModel)

	if tmpUser, err = userModel.GetUser(mobile); err != nil {
		err = libs.NewReportError(err)
	}

	if tmpUser.Id > 0 {
		return tmpUser, libs.NewReportError(errors.New("The phone has already to register. "))
	}

	// 插入数据
	userModel.Avatar = avatar
	userModel.Mobile = mobile
	userModel.Nickname = nickName
	userModel.Sex = sex
	// 加密
	userModel.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	userModel.Passwd = utils.MakePasswd(plainPwd, userModel.Salt)

	_, err = userModel.Insert()

	if err != nil {
		err = libs.NewReportError(err)
	}

	user = userModel

	return

}

// 登录函数
func (s *UserService) Login(mobile, plainPwd string) (user model.UserModel, err error) {

	return

}
