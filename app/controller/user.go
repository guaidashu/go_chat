package controller

import (
	"github.com/gin-gonic/gin"
	"go_chat/app/libs"
)

func UserLogin(ctx *gin.Context) {
	libs.Success(ctx, "user login")
}
