package types

import (
	"web_server_gin/pkg/model"
)

type PostCreate struct {
	model.Post
}

type PostResponse struct {
	// model.Post
	// Name        string    `json:"name" gorm:"size:64"`
	// Description string    `json:"description" gorm:"size:256;null" binding:"lt=256"`
	ID        int64  `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;unique;Column:id"`
	CreatedAt string `json:"created_at" gorm:"not null" time_format:"2006-01-02"`
	UpdatedAt string `json:"updated_at" gorm:"null" time_format:"2006-01-02"`
	Creator   string `json:"creator" gorm:"size:32" binding:"lt=32"`
	Modifier  string `json:"modifier" gorm:"size:32;null" binding:"lt=32"`
	Title     string `json:"title" gorm:"size:128;not null;unique" binding:"required,lt=128"`
	Content   string `json:"content" gorm:"type:text;not null" binding:"required"`
	Author    string `json:"author" gorm:"size:32;not null" binding:"lt=32"`
	Category  string `json:"category" gorm:"size:32;not null" binding:"required,lt=32"`
	Read      int64  `json:"read" binding:"max=10,min=0"`
	Comment   int64  `json:"comment" binding:"max=10,min=0"`
}

// PostsResponse ...
type PostsResponse struct {
	Count int64          `json:"count"`
	Posts []PostResponse `json:"posts"`
}

// RenderPostResp ...
func RenderPostResp(post model.Post) PostResponse {
	PostResp := PostResponse{
		post.ID, post.CreatedAt.Format("2006-01-02 15:04:05"), post.UpdatedAt.Format("2006-01-02 15:04:05"),
		post.Creator, post.Modifier, post.Title,
		post.Content, post.Author, post.Category,
		post.Read, post.Comment,
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
