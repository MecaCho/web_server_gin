package model

// {{.title}}, {{.content}}, {{.comment}}, {{.author}}, {{.read}}, {{category}} {{.created_at}} {{.id}}
type Post struct {
	// gorm.Model
	BaseCommonModel
	Title    string `json:"title" gorm:"size:128;not null;unique" binding:"required"`
	Content  string `json:"content" gorm:"type:text;not null" binding:"required"`
	Author   string `json:"author" gorm:"size:32;not null"`
	Category string `json:"category" gorm:"size:32;not null" binding:"required"`
	Read     int64  `json:"read"`
	Comment  int64  `json:"comment"`
}

func (p *Post) TableName() string {
	return DBTableNamePost
}
