/**
  create by yy on 2019-08-23
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"go_chat/app/libs"
)

func Test(ctx *gin.Context) {
	//user := new(models.UserModel)
	//data, _ := user.GetUserById(1)
	//libs.Logger.Errorf("ok")
	libs.Success(ctx, "test is ok")
}
