/**
  create by yy on 2019-08-23
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/index/index.html", gin.H{})
}
