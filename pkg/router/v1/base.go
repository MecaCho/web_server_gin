package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"web_server_gin/pkg/controller"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/model"
)

func RouterGroup(dbCon string, router *gin.Engine) (err error) {

	db, err := gorm.Open("mysql", dbCon)
	if err != nil {
		return
	}
	db.AutoMigrate(&model.Post{})
	dbORM := dao.DB{*db, true}
	server := controller.ServerHandle{dbORM}

	// router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/posts", server.CreateResourceController)
		v1.PUT("/posts/:post_id", server.UpdateResourceController)
		v1.GET("/posts", server.ListResourcesController)
		v1.GET("/posts/:post_id", server.GetResourceController)
		v1.DELETE("/posts/:post_id", server.DeleteResourcesController)
	}
	return
}
