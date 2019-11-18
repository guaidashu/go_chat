package service

import (
	"errors"
	"fmt"
	"github.com/guaidashu/go_helper/crypto_tool"
	"go_chat/app/libs"
	"go_chat/origin/model"
	"go_chat/origin/utils"
	"math/rand"
	"time"
)

type UserService struct {
}

// 注册函数
func (s *UserService) Register(mobile, plainPwd, nickName, avatar, sex string) (user *model.UserModel, err error) {

	var tmpUser *model.UserModel

	userModel := new(model.UserModel)

	if tmpUser, err = userModel.GetUser(mobile); err != nil {
		err = libs.NewReportError(err)
		return
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

	userModel.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = userModel.Insert()

	if err != nil {
		err = libs.NewReportError(err)
	}

	user = userModel

	return

}

// 登录函数
func (s *UserService) Login(mobile, plainPwd string) (user *model.UserModel, err error) {

	userModel := new(model.UserModel)

	if user, err = userModel.GetUser(mobile); err != nil {
		err = libs.NewReportError(err)
		return
	}

	if user.Id == 0 {
		err = libs.NewReportError(errors.New("user is not exists or password is error"))
		return
	}

	if !utils.ValidatePasswd(plainPwd, user.Salt, user.Passwd) {
		err = libs.NewReportError(errors.New("user is not exists or password is error"))
		return
	}

	// 刷新token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := crypto_tool.MD5(str)
	user.Token = token

	// 刷新数据库的token
	if _, err = user.UpdateToken(); err != nil {
		err = libs.NewReportError(err)
	}

	return

}
