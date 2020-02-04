package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"time"
	"web_server_gin/config"
	"web_server_gin/html"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/router/v1"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()
		glog.Infof(
			"%s %s %s %s %d",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL,
			latency,
			status,
		)
		glog.Infof("Request remote addr: %s.", c.Request.RemoteAddr)
	}
}

func main() {
	engine := gin.Default()
	engine.Use(Logger())
	engine.StaticFS("/static", http.Dir("./html/static"))
	engine.LoadHTMLGlob("html/html/*")
	// engine.Any("/", controller.HelloController)
	conf := config.InitConfig()
	dbORM, err := dao.InitDB(conf.DBConn)
	if err != nil {
		panic(err)
	}

	if err := html.HTMLRouter(dbORM, engine); err != nil {
		panic(err)
	}

	if err := v1.RouterGroup(dbORM, engine); err != nil {
		panic(err)
	}

	engine.Run(fmt.Sprintf("%s:%d", conf.ServerIP, conf.ServerPort))
}
