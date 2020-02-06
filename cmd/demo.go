package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"time"
	"web_server_gin/config"
	"web_server_gin/html"
	"web_server_gin/pkg/common"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/router/v1"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("Content-Type", "application/json")

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
	}
}

func main() {
	engine := gin.Default()
	binding.Validator = common.NewStructValidator
	// binding.Validator.Engine() = binding.Validator.Engine()
	// binding.Validator.RegisterValidation("timevalidate", common.TimeValidate)
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
