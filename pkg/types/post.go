package types

import "web_server_gin/pkg/model"

type PostCreate struct {
	model.Post
}

type PostResponse struct {
	model.Post
}

// PostsResponse ...
type PostsResponse struct {
	Count int64          `json:"count"`
	Posts []PostResponse `json:"posts"`
}

// RenderPostResp ...
func RenderPostResp(post model.Post) PostResponse {
	PostResp := PostResponse{
		post,
	}
	return PostResp
}

// AddResp ...
func (tr *PostsResponse) AddResp(Posts []model.Post) {
	for _, post := range Posts {
		tr.Posts = append(tr.Posts, RenderPostResp(post))
	}
}

// NewPostsResponse ...
func NewPostsResponse(num int64, posts []model.Post) PostsResponse {
	var PostResp PostsResponse
	PostResp.Count = num
	PostResp.AddResp(posts)
	if PostResp.Count == 0 {
		PostResp.Posts = []PostResponse{}
	}
	return PostResp
}
