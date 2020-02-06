package model

type Comment struct {
	BaseCommonModel
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Content  string `json:"content"`
	PostID   int64  `json:"post_id"`
}
