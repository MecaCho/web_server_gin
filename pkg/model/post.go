package model

// {{.title}}, {{.content}}, {{.comment}}, {{.author}}, {{.read}}, {{category}} {{.created_at}} {{.id}}
type Post struct {
	// gorm.Model
	BaseCommonModel
	Title    string `json:"title" gorm:"size:128;not null;unique" binding:"required,lt=128"`
	Content  string `json:"content" gorm:"type:text;not null" binding:"required"`
	Author   string `json:"author" gorm:"size:32;not null" binding:"lt=32"`
	Category string `json:"category" gorm:"size:32;not null" binding:"required,lt=32"`
	Read     int64  `json:"read" binding:"max=10,min=0"`
	Comment  int64  `json:"comment" binding:"max=10,min=0"`
}

func (p *Post) TableName() string {
	return DBTableNamePost
}
