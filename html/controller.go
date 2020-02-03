package html

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/model"
)

type ServerHandle struct {
	ORM *dao.DB
}

func HTMLRouter(dbORM *dao.DB, router *gin.Engine) (err error) {
	server := ServerHandle{dbORM}
	v1 := router.Group("/")
	{
		v1.GET("", server.IndexController)
		v1.GET("index.html", server.IndexController)
		v1.GET("contact.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "contact.html", "hello, I am qiuwenqi.")
		})
		v1.GET("about.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "about.html", "hello, I am qiuwenqi.")
		})
		v1.GET("/posts/:post_id", server.GetPostController)
		v1.GET("full-width.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "full-width.html", "hello, I am qiuwenqi.")
		})
	}
	return
}

func (sh *ServerHandle) IndexController(ctx *gin.Context) {
	var posts []model.Post
	filters := ctx.Request.URL.Query()
	glog.Infof("query filters :%+v", filters)
	_, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	if err != nil {
		glog.Errorf("List posts error: %s.", err.Error())
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data": posts,
	})
}

// GetResourceController ...
func (sh *ServerHandle) GetPostController(ctx *gin.Context) {
	var post model.Post
	filters := map[string][]string{}

	id, _ := ctx.Params.Get("post_id")
	// if !ok {
	// 	ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, "resource not found"))
	// 	return
	// }

	filters["id"] = []string{id}
	sh.ORM.FilterTable(filters, &post, dao.DBTableNamePost)
	// if err != nil || num == 0 {
	// 	ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, err.Error()))
	// 	return
	// }
	var posts []model.Post
	posts = append(posts, post)
	ctx.HTML(http.StatusOK, "single.html", gin.H{
		"posts": posts,
	})

	return
}
