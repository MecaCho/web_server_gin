package dao

import (
	"github.com/golang/glog"
	"web_server_gin/pkg/model"
)

func (db *DB) UpdatePost(post model.Post) (err error) {
	err = db.ORM.Model(post).Updates(post).Error
	if err != nil {
		glog.Errorf("update resource error: %s.", err.Error())
	}
	return
}
