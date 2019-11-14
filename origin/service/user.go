package service

import "go_chat/origin/model"

type UserService struct {
}

// 注册函数
func (s *UserService) Register(mobile, plainPwd, nickName, avatar, sex string) (user model.UserModel, err error) {



	return

}

// 登录函数
func (s *UserService) Login(mobile, plainPwd string) (user model.UserModel, err error) {

	return

}