package dao

import (
	"github.com/jinzhu/gorm"
	"web_server_gin/pkg/model"
)

func InitDB(dbCon string) (dbORM *DB, err error) {
	db, err := gorm.Open("mysql", dbCon)
	if err != nil {
		return
	}
	db.AutoMigrate(&model.Post{}, &model.Comment{}, &model.PersonaInfo{})
	dbORM = &DB{*db, true}
	return
}
