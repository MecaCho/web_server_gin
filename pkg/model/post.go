package model

import "time"

// {{.title}}, {{.content}}, {{.comment}}, {{.author}}, {{.read}}, {{category}} {{.created_at}} {{.id}}
type Post struct {
	// gorm.Model
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;unique;Column:id"`
	Name        string    `json:"name" gorm:"size:64"`
	Description string    `json:"description" gorm:"size:256;null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"null"`
	Creator     string    `json:"creator" gorm:"size:32"`
	Modifier    string    `json:"modifier" gorm:"size:32;null"`
	Title       string    `json:"title" gorm:"size:128"`
	Content     string    `json:"content" gorm:"type:text"`
	Author      string    `json:"author" gorm:"size:32"`
	Category    string    `json:"category"`
	Read        int64     `json:"read"`
	Comment     int64     `json:"comment"`
}

func (p *Post) TableName() string {
	return DBTableNamePost
}
