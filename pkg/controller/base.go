package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_server_gin/pkg/dao"
)

type ServerHandle struct {
	ORM dao.DB
}

func HelloController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, I am qiuwenqi.")
}
