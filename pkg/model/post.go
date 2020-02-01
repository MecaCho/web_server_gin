package model

type Post struct {
	// gorm.Model
	BaseCommonModel
	Title   string `json:"title" orm:"size:128"`
	Content string `json:"content" orm:"type:text"`
}
