package html

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HTMLController(router *gin.Engine) (err error) {
	v1 := router.Group("/")
	{
		v1.GET("index.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", "hello, I am qiuwenqi.")
		})
		v1.GET("contact.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "contact.html", "hello, I am qiuwenqi.")
		})
		v1.GET("about.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "about.html", "hello, I am qiuwenqi.")
		})
		v1.GET("single.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "single.html", "hello, I am qiuwenqi.")
		})
		v1.GET("full-width.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "full-width.html", "hello, I am qiuwenqi.")
		})
	}
	return
}
