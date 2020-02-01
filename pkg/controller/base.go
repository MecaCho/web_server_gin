package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloController(ctx *gin.Context)  {
	ctx.String(http.StatusOK, "hello, I am qiuwenqi.")
}
