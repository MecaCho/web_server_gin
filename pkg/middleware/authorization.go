package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"web_server_gin/pkg/types"
)

func CheckAuthorization(ctx *gin.Context) (err error) {
	user := ctx.PostForm("username")
	password := ctx.PostForm("password")

	glog.Infof("request body: %s, %s.", user, password)
	if user != "qwq" || password != "qwq" {
		return types.NewErrorResponse(500, "unauthorized")
	}
	return
}

func AuthMiddleWare(ctx *gin.Context) {
	if err := CheckAuthorization(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		ctx.Abort()
	}
	ctx.Next()
}
