package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"web_server_gin/config"
	"web_server_gin/pkg/controller"
	"web_server_gin/pkg/router/v1"
)

func main() {
	engine := gin.Default()
	engine.Any("/", controller.HelloController)
	conf := config.InitConfig()
	v1.RouterGroup(conf.DBConn, engine)

	engine.Run(":80")
}
