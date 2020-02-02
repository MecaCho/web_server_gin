package model

import "web_server_gin/pkg/dao"

type Post struct {
	// gorm.Model
	BaseCommonModel
	Title   string `json:"title" orm:"size:128"`
	Content string `json:"content" orm:"type:text"`
}

func (p *Post) TableName() string {
	return dao.DBTableNamePost
}
