package html

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		v1.GET("about.html", server.GetPersonalInfoController)
		v1.GET("/posts/:post_id", server.GetPostController)
		v1.POST("/posts/:post_id", server.CreatePostCommentController)
		v1.GET("full-width.html", server.IndexController)

		v1.GET("add.html", func(context *gin.Context) {
			if err := middleware.CheckAuthorization(context); err != nil {
				context.JSON(http.StatusUnauthorized, err)
				return
			}
			context.HTML(http.StatusOK, "add.html", "hello, I am qiuwenqi.")
		})
		v1.GET("posts", server.IndexController)
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

func (sh *ServerHandle) GetPersonalInfoController(ctx *gin.Context) {
	var info model.PersonaInfo
	_, info, err := sh.ORM.GetPersonalInfo(nil)
	if err != nil{
		glog.Errorf("Get personal info error: %s.", err.Error())
	}
	// glog.Infof("get personal info: %+v.", info)

	ctx.HTML(http.StatusOK, "about.html", gin.H{
		"introduction": info.Introduction,
		"facebook":     info.Facebook,
		"twitter":      info.Twitter,
		"zhihu":        info.Zhihu,
		"weibo":        info.Weibo,
		"wechat":       info.Wechat,
	})

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
	tag := ctx.PostForm("tag")

	postCreate.Content = content
	postCreate.Title = title
	postCreate.Category = category
	postCreate.Author = author
	if tag == "" {
		postCreate.Tag = title[:4]
	}
	postCreate.Archive = time.Now().Format("2006-01")

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

	postDetail, contentRender, postsRender, err := sh.GetPostDetail(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}
	ctx.HTML(http.StatusOK, "single.html", gin.H{
		"comment_num":    postDetail.Comment,
		"comments":       postDetail.Comments,
		"content_render": contentRender.ContentRender,
		"post_id":        postDetail.ID,
		"title":          postDetail.Title,
		"read":           postDetail.Read,
		"comment":        postDetail.Comment,
		"category":       postDetail.Category,
		"author":         postDetail.Author,
		"created_at":     postDetail.CreatedAt,
		"categorys":      postsRender.Categories,
		"tags":           postsRender.Tags,
		"archives":       postsRender.Archives,
		"newest_posts":   postsRender.NewestPosts,
	})
	return
}

type PostsRenderResponse struct {
	PostsResponses types.PostsResponse `json:"posts_responses"`
	Archives       []Archive           `json:"archives"`
	Categories     []Category          `json:"category"`
	Tags           []Tag               `json:"tags"`
	NewestPosts    []NewPost           `json:"newest_posts"`
}

func ConvertPostsRender(posts []model.Post, shortCut bool) (postsRender PostsRenderResponse) {
	var categorys []Category
	var archives []Archive
	categroyMap := make(map[string]int64)
	archivesMap := make(map[string]int64)
	tagMap := make(map[string]int64)
	var tags []Tag

	var newestPosts []NewPost

	for k, post := range posts {
		glog.Infof("content short cut: %+v.", shortCut)
		if shortCut == true {
			posts[k].Content = PrintContentShortCut(post.Content)
		}
		categoryKey := post.Category
		if _, ok := categroyMap[categoryKey]; ok {
			categroyMap[categoryKey] += 1
		} else {
			categroyMap[categoryKey] = 1
		}

		archivesKey := post.CreatedAt.Format("2006-01")
		if _, ok := archivesMap[archivesKey]; ok {
			archivesMap[archivesKey] += 1
		} else {
			archivesMap[archivesKey] = 1
		}

		tagName := post.Title
		if _, ok := tagMap[tagName]; ok {
			tagMap[tagName] += 1
		} else {
			tagMap[tagName] = 1
		}
		if len(newestPosts) < 7 {
			newestPosts = append(newestPosts, NewPost{post.ID, post.Title})
		}
	}

	for k, value := range categroyMap {
		categorys = append(categorys, Category{k, value})
	}
	for k, value := range archivesMap {
		archives = append(archives, Archive{k, value})
	}
	for k, _ := range tagMap {
		tags = append(tags, Tag{k})
	}
	postsResponse := types.NewPostsResponse(int64(len(posts)), posts)

	postsRender.Archives = archives
	postsRender.Categories = categorys
	postsRender.Tags = tags
	postsRender.NewestPosts = newestPosts
	postsRender.PostsResponses = postsResponse
	return
}

func PrintContentShortCut(content string) string {
	if len(content) > 128 {
		lines := strings.Split(content, "\r")
		if strings.Contains(content, "\n") {
			lines = strings.Split(content, "\n")
		}
		glog.Infof("get content first line: %+v.", len(lines[0]))
		if len(lines[0]) > 128 {
			return lines[0][:128]
		}
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

	postsRender := ConvertPostsRender(posts, true)
	var pages []Page
	count := postsRender.PostsResponses.Count
	for i := 0; i <= int(count/21); i++ {
		pages = append(pages, Page{i + 1, i * 20})
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data":         postsRender.PostsResponses.Posts,
		"count":        postsRender.PostsResponses.Count,
		"categorys":    postsRender.Categories,
		"tags":         postsRender.Tags,
		"newest_posts": postsRender.NewestPosts,
		"archives":     postsRender.Archives,
		"pages":        pages,
	})
}

type Page struct {
	PageNum int `json:"page_num"`
	Offset  int `json:"offset"`
}

type PostRender struct {
	ContentRender template.HTML `json:"content_render"`
}

type Category struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type Archive struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}

type Tag struct {
	Name string `json:"name"`
}

type NewPost struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (sh *ServerHandle) GetPostDetail(id string) (postDetail types.PostResponse, contentRender PostRender, postsRender PostsRenderResponse, err error) {
	filters := map[string][]string{}
	filters["id"] = []string{id}
	num, posts, err := sh.ORM.ListPosts(filters)
	if err != nil || num == 0 {
		err = types.NewErrorResponse(common.NotFound, err.Error())
		return
	}
	postsResponse := types.NewPostsResponse(int64(len(posts)), posts)
	postDetail = postsResponse.Posts[0]
	// var postRender PostRender
	switch {
	case strings.Contains(postDetail.Content, "\n"):
		contentRender.ContentRender = template.HTML(strings.Join(strings.Split(postDetail.Content, "\n"), "<br />"))
	case strings.Contains(postDetail.Content, "\r"):
		contentRender.ContentRender = template.HTML(strings.Join(strings.Split(postDetail.Content, "\r"), "<br />"))
	default:
		contentRender.ContentRender = template.HTML(strings.Join(strings.Split(postDetail.Content, "\t"), "<br />"))
	}
	// for k, post := range postsResponse.Posts {
	// 	glog.Infof("Post content: %s.", postsResponse.Posts[k].Content)
	// 	content := post.Content
	// 	postRender = template.HTML(strings.Join(strings.Split(content, "\n"), "<br />")
	// 	// postsResponse.Posts[k].Content = strings.ReplaceAll(post.Content, "\n", "<br/>")
	// 	glog.Infof("Post content: %s.", postsResponse.Posts[k].Content)
	// }
	postsRender = ConvertPostsRender(posts, false)
	return
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
	postDetail, contentRender, postsRender, err := sh.GetPostDetail(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}

	glog.Infof("post comment num: %d.", len(postDetail.Comments))
	ctx.HTML(http.StatusOK, "single.html", gin.H{
		"comment_num":    postDetail.Comment,
		"comments":       postDetail.Comments,
		"content_render": contentRender.ContentRender,
		"post_id":        postDetail.ID,
		"title":          postDetail.Title,
		"read":           postDetail.Read,
		"comment":        postDetail.Comment,
		"category":       postDetail.Category,
		"author":         postDetail.Author,
		"created_at":     postDetail.CreatedAt,
		"categorys":      postsRender.Categories,
		"tags":           postsRender.Tags,
		"archives":       postsRender.Archives,
		"newest_posts":   postsRender.NewestPosts,
	})

	posts[0].Read += 1
	sh.ORM.UpdatePostColumn(posts[0], "read", strconv.Itoa(int(posts[0].Read)))
	return
}
