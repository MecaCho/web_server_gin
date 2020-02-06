package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/model"
	"web_server_gin/pkg/types"
)

// ListPostCommentsController ...
func (sh *ServerHandle) ListPostCommentsController(ctx *gin.Context) {
	// filters := ctx.Request.URL.Query()
	// glog.Infof("query filters :%+v", filters)
	var posts []model.Post
	filters := map[string][]string{}

	id, ok := ctx.Params.Get("post_id")
	if !ok {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, "resource not found"))
		return
	}

	filters["id"] = []string{id}
	num, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	if err != nil || num == 0 {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, err.Error()))
		return
	}
	for k, _ := range posts {
		sh.ORM.RelodeComment(&posts[k])
	}
	ctx.JSON(http.StatusOK, types.RenderPostResp(posts[0]))
	sh.ORM.UpdatePost(posts[0])
	return
}

// GetPostCommentController ...
func (sh *ServerHandle) GetPostCommentController(ctx *gin.Context) {
	var posts []model.Post
	filters := map[string][]string{}

	id, ok := ctx.Params.Get("post_id")
	if !ok {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, "resource not found"))
		return
	}

	filters["id"] = []string{id}
	num, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	if err != nil || num == 0 {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, err.Error()))
		return
	}
	for k, _ := range posts {
		sh.ORM.RelodeComment(&posts[k])
	}
	ctx.JSON(http.StatusOK, types.RenderPostResp(posts[0]))
	sh.ORM.UpdatePost(posts[0])
	return
}

// UpdatePostCommentController ...
func (sh *ServerHandle) UpdatePostCommentController(ctx *gin.Context) {
	return
}

// DeletePostCommentsController ...
func (sh *ServerHandle) DeletePostCommentsController(ctx *gin.Context) {
	return
}

// CreatePostCommentController ...
func (sh *ServerHandle) CreatePostCommentController(ctx *gin.Context) {
	var posts []model.Post
	filters := map[string][]string{}

	id, ok := ctx.Params.Get("post_id")
	if !ok {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, "resource not found"))
		return
	}

	filters["id"] = []string{id}
	num, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	if err != nil || num == 0 {
		ctx.JSON(http.StatusNotFound, types.NewErrorResponse(500, err.Error()))
		return
	}
	for k, _ := range posts {
		sh.ORM.RelodeComment(&posts[k])
	}

	var comment model.Comment
	if err := ctx.ShouldBind(&comment); err != nil {
		glog.Errorf("Validate binding request body error: %s.", err.Error())
		ctx.JSON(http.StatusBadRequest, types.NewErrorResponse(400, err.Error()))
		return
	}
	comment.PostID = posts[0].ID
	// err := binding.Validator.ValidateStruct(postCreate.Post)
	// if err != nil {
	// 	glog.Errorf("Validate request body error: %s.", err.Error())
	// 	ctx.JSON(http.StatusBadRequest, types.NewErrorResponse(400, err.Error()))
	// 	return
	// }
	err = sh.ORM.CreateResource(&comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, comment)
	return
}
