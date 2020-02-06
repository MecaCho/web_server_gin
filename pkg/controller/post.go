package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"web_server_gin/pkg/dao"
	"web_server_gin/pkg/model"
	"web_server_gin/pkg/types"
)

func CreatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusCreated, "Hello %s", id)
}

// ListResourcesController ...
func (sh *ServerHandle) ListResourcesController(ctx *gin.Context) {
	var posts []model.Post
	filters := ctx.Request.URL.Query()
	glog.Infof("query filters :%+v", filters)

	num, err := sh.ORM.FilterTable(filters, &posts, dao.DBTableNamePost)
	// num = int64(len(posts))
	glog.Infof("query result num:%d, %d.", num, len(posts))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, types.NewPostsResponse(num, posts))
	return
}

// GetResourceController ...
func (sh *ServerHandle) GetResourceController(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, types.RenderPostResp(posts[0]))
	return
}

// UpdateResourceController ...
func (sh *ServerHandle) UpdateResourceController(ctx *gin.Context) {
	return
}

// DeleteResourcesController ...
func (sh *ServerHandle) DeleteResourcesController(ctx *gin.Context) {
	return
}

// CreateResourceController ...
func (sh *ServerHandle) CreateResourceController(ctx *gin.Context) {
	var postCreate types.PostCreate
	if err := ctx.ShouldBind(&postCreate.Post); err != nil {
		glog.Errorf("Validate binding request body error: %s.", err.Error())
		ctx.JSON(http.StatusBadRequest, types.NewErrorResponse(400, err.Error()))
		return
	}
	// err := binding.Validator.ValidateStruct(postCreate.Post)
	// if err != nil {
	// 	glog.Errorf("Validate request body error: %s.", err.Error())
	// 	ctx.JSON(http.StatusBadRequest, types.NewErrorResponse(400, err.Error()))
	// 	return
	// }
	err := sh.ORM.CreateResource(&postCreate.Post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewErrorResponse(500, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, types.RenderPostResp(postCreate.Post))
	return
}
