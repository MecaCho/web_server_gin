package model

type Comment struct {
	BaseCommonModel
	UserName string `json:"user_name" binding:"required,lt=128"`
	Email    string `json:"email" binding:"required,lt=128"`
	Content  string `json:"content" binding:"required,lt=128"`
	PostID   int64  `json:"post_id"`
}
