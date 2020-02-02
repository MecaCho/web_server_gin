package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"web_server_gin/config"
	"web_server_gin/pkg/controller"
	"web_server_gin/pkg/router/v1"
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

func main() {
	engine := gin.Default()
	engine.StaticFS("/static", http.Dir("/Users/rainmc/GO/src/web_server_gin/html/static"))
	// engine.StaticFS("/*.html", http.Dir("/Users/rainmc/GO/src/web_server_gin/html/html"))
	engine.LoadHTMLGlob("html/html/*")
	engine.Any("/", controller.HelloController)
	HTMLController(engine)
	conf := config.InitConfig()

	if err := v1.RouterGroup(conf.DBConn, engine); err != nil {
		panic(err)
	}

	engine.Run(fmt.Sprintf("%s:%d", conf.ServerIP, conf.ServerPort))
}
