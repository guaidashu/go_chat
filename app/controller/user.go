package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/user/login.shtml", gin.H{})
}

func UserRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/user/register.shtml", gin.H{})
}
