package main

import (
	"github.com/gin-gonic/gin"
	"web_server_gin/pkg/controller"
)

func main()  {
	engine := gin.Default()
	engine.Any("/", controller.HelloController)

	engine.Run(":80")
}
