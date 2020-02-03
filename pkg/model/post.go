package model

// {{.title}}, {{.content}}, {{.comment}}, {{.author}}, {{.read}}, {{category}} {{.created_at}} {{.id}}
type Post struct {
	// gorm.Model
	BaseCommonModel
	Title    string `json:"title" gorm:"size:128"`
	Content  string `json:"content" gorm:"type:text"`
	Author   string `json:"author" gorm:"size:32"`
	Category string `json:"category"`
	Read     int64  `json:"read"`
	Comment  int64  `json:"comment"`
}

func (p *Post) TableName() string {
	return DBTableNamePost
}
