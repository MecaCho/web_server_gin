package html

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"net/http"
	"web_server_gin/pkg/common"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/model"
	"web_server_gin/pkg/types"
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
		v1.GET("full-width.html", server.IndexController)

		v1.GET("add.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "add.html", "hello, I am qiuwenqi.")
		})
		v1.POST("posts", server.CreatePostController)
		v1.GET("admin.html", func(context *gin.Context) {
			context.HTML(http.StatusOK, "admin.html", "hello, I am qiuwenqi.")
		})
	}
	return
}
func (sh *ServerHandle) CreatePostController(ctx *gin.Context) {
	var postCreate types.PostCreate

	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	category := ctx.PostForm("category")
	author := ctx.PostForm("author")

	postCreate.Content = content
	postCreate.Title = title
	postCreate.Category = category
	postCreate.Author = author

	err := binding.Validator.ValidateStruct(postCreate.Post)
	if err != nil {
		glog.Errorf("Validate request body error: %s.", err.Error())
		ctx.JSON(http.StatusBadRequest, types.NewErrorResponse(400, err.Error()))
		return
	}

	err = sh.ORM.CreateResource(&postCreate.Post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}
	var posts []model.Post
	posts = append(posts, postCreate.Post)
	ctx.HTML(http.StatusCreated, "single.html", gin.H{
		"posts": posts,
	})
	return

}

func PrintContentShortCut(content string) string {
	if len(content) > 128 {
		return content[:128]
	} else {
		return content
	}
}

func (sh *ServerHandle) IndexController(ctx *gin.Context) {
	var posts []model.Post
	filters := ctx.Request.URL.Query()
	glog.Infof("query filters :%+v", filters)
	_, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	if err != nil {
		glog.Errorf("List posts error: %s.", err.Error())
	}

	for k, post := range posts {
		posts[k].Content = PrintContentShortCut(post.Content)
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data": posts,
	})
}

// GetResourceController ...
func (sh *ServerHandle) GetPostController(ctx *gin.Context) {
	var posts []model.Post
	filters := map[string][]string{}

	id, _ := ctx.Params.Get("post_id")
	// if !ok {
	// 	ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, "resource not found"))
	// 	return
	// }

	filters["id"] = []string{id}
	num, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	if err != nil || num == 0 {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(common.NotFound, err.Error()))
		return
	}
	ctx.HTML(http.StatusOK, "single.html", gin.H{
		"posts": posts,
	})

	return
}
