package html

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"web_server_gin/pkg/common"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/middleware"
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
		v1.POST("/posts/:post_id/comments", server.CreatePostCommentController)
		v1.GET("full-width.html", server.IndexController)

		v1.GET("add.html", func(context *gin.Context) {
			if err := middleware.CheckAuthorization(context); err != nil {
				context.JSON(http.StatusUnauthorized, err)
				return
			}
			context.HTML(http.StatusOK, "add.html", "hello, I am qiuwenqi.")
		})
		v1.POST("posts", server.CreatePostController)
	}

	admin := router.Group("/admin")
	{
		admin.GET("", func(context *gin.Context) {
			context.HTML(http.StatusOK, "login.html", "")
		})
		admin.GET("login", func(context *gin.Context) {
			context.HTML(http.StatusOK, "login.html", "")
		})
		admin.POST("login", server.LoginController)
	}
	return
}

func (sh *ServerHandle) LoginController(ctx *gin.Context) {
	if err := middleware.CheckAuthorization(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.HTML(http.StatusCreated, "add.html", gin.H{})
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

func (sh *ServerHandle) CreatePostCommentController(ctx *gin.Context) {
	var comment model.Comment

	postID, _ := ctx.Params.Get("post_id")
	// filters := make(map[string][]string)
	// _, posts, _ := sh.ORM.ListPosts(filters)
	// post := posts[0].ID

	user_name := ctx.PostForm("user_name")
	content := ctx.PostForm("content")
	email := ctx.PostForm("email")

	comment.UserName = user_name
	comment.Content = content
	comment.Email = email
	postIDInt, _ := strconv.Atoi(postID)
	comment.PostID = int64(postIDInt)

	err := binding.Validator.ValidateStruct(comment)
	if err != nil {
		glog.Errorf("Validate request body error: %s.", err.Error())
		ctx.JSON(http.StatusBadRequest, types.NewErrorResponse(400, err.Error()))
		return
	}

	err = sh.ORM.CreateResource(&comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}
	return
}

func PrintContentShortCut(content string) string {
	if len(content) > 128 {
		lines := strings.Split(content, "\n")
		glog.Infof("get content first line: %s.", lines[0])
		return lines[0]
	} else {
		return content
	}
}

func (sh *ServerHandle) IndexController(ctx *gin.Context) {
	var posts []model.Post
	filters := ctx.Request.URL.Query()
	glog.Infof("query filters :%+v", filters)
	_, posts, err := sh.ORM.ListPosts(filters)
	if err != nil {
		glog.Errorf("List posts error: %s.", err.Error())
	}

	var categorys []Category
	categroyMap := make(map[string]int64)
	tagMap := make(map[string]int64)
	var tags []Tag

	for k, post := range posts {
		posts[k].Content = PrintContentShortCut(post.Content)
		categoryKey := post.Category
		if _, ok := categroyMap[categoryKey]; ok {
			categroyMap[categoryKey] += 1
		} else {
			categroyMap[categoryKey] = 1
		}
		tagName := post.Title
		if _, ok := tagMap[tagName]; ok {
			tagMap[tagName] += 1
		} else {
			tagMap[tagName] = 1
		}
	}
	for k, value := range categroyMap {
		categorys = append(categorys, Category{k, value})
	}
	for k, _ := range tagMap {
		tags = append(tags, Tag{k})
	}
	postsResponse := types.NewPostsResponse(int64(len(posts)), posts)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data":      postsResponse.Posts,
		"categorys": categorys,
		"tags":      tags,
	})
}

type PostRender struct {
	ContentRender template.HTML `json:"content_render"`
}
type Category struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
type Tag struct {
	Name string `json:"name"`
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
	num, posts, err := sh.ORM.ListPosts(filters)
	if err != nil || num == 0 {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(common.NotFound, err.Error()))
		return
	}
	postsResponse := types.NewPostsResponse(int64(len(posts)), posts)
	postDetail := postsResponse.Posts[0]
	var postRender PostRender
	postRender.ContentRender = template.HTML(strings.Join(strings.Split(postDetail.Content, "\n"), "<br />"))
	// for k, post := range postsResponse.Posts {
	// 	glog.Infof("Post content: %s.", postsResponse.Posts[k].Content)
	// 	content := post.Content
	// 	postRender = template.HTML(strings.Join(strings.Split(content, "\n"), "<br />")
	// 	// postsResponse.Posts[k].Content = strings.ReplaceAll(post.Content, "\n", "<br/>")
	// 	glog.Infof("Post content: %s.", postsResponse.Posts[k].Content)
	// }
	var categorys []Category
	var tags []Tag

	glog.Infof("post comment num: %d.", len(postDetail.Comments))
	ctx.HTML(http.StatusOK, "single.html", gin.H{
		"posts":          postsResponse.Posts,
		"comment_num":    postDetail.Comment,
		"comments":       postDetail.Comments,
		"content_render": postRender.ContentRender,
		"post_id":        postDetail.ID,
		"title":          postDetail.Title,
		"read":           postDetail.Read,
		"comment":        postDetail.Comment,
		"category":       postDetail.Category,
		"author":         postDetail.Author,
		"created_at":     postDetail.CreatedAt,
		"categorys":      categorys,
		"tags":           tags,
	})

	posts[0].Read += 1
	sh.ORM.UpdatePost(posts[0])
	return
}
