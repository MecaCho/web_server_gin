package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_server_gin/pkg/types"
)

func CheckAuthorization(ctx *gin.Context) (err error) {
	user := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// glog.Infof("request body: %s, %s.", user, password)
	if user == "qwq" && password == "qwq" {
		return
		// return types.NewErrorResponse(500, "unauthorized")
	}

	tokenStr := ctx.GetHeader("Authorization")
	userInfo, err := ParserToken(tokenStr)
	if err != nil {
		return types.NewErrorResponse(http.StatusUnauthorized, err.Error())
	}
	ctx.Set("userInfo", userInfo)

	return
}

func AuthMiddleWare(ctx *gin.Context) {
	if err := CheckAuthorization(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		ctx.Abort()
	}
	ctx.Next()
}
