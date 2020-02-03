package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"web_server_gin/config"
	"web_server_gin/html"
	"web_server_gin/pkg/controller"
	"web_server_gin/pkg/router/v1"
)

func main() {
	engine := gin.Default()
	engine.StaticFS("/static", http.Dir("./html/static"))
	engine.LoadHTMLGlob("html/html/*")
	engine.Any("/", controller.HelloController)
	html.HTMLController(engine)
	conf := config.InitConfig()

	if err := v1.RouterGroup(conf.DBConn, engine); err != nil {
		panic(err)
	}

	engine.Run(fmt.Sprintf("%s:%d", conf.ServerIP, conf.ServerPort))
}
