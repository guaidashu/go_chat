package controller

import (
	"github.com/gin-gonic/gin"
	"go_chat/app/libs"
	"go_chat/app/models"
)

func InitTables(ctx *gin.Context) {
	models.CreateTable()
	libs.Success(ctx, "create table")
}
