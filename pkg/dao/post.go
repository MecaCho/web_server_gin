package dao

import (
	"github.com/golang/glog"
	"web_server_gin/pkg/model"
)

func (db *DB) UpdatePost(post model.Post) (err error) {
	if post.Archive == "" {
		post.Archive = post.CreatedAt.Format("2006-01")
	}
	err = db.ORM.Model(post).Updates(post).Error
	if err != nil {
		glog.Errorf("update resource error: %s.", err.Error())
	}
	return
}

func (db *DB) UpdatePostColumn(post model.Post, columnName, value string) (err error) {
	err = db.ORM.Model(post).UpdateColumn(columnName, value).Error
	if err != nil {
		glog.Errorf("update resource error: %s.", err.Error())
	}
	if post.Archive == "" {
		post.Archive = post.CreatedAt.Format("2006-01")
		db.ORM.Model(post).UpdateColumn("archive", post.Archive)
	}
	return

}

func (db *DB) ListPosts(filters map[string][]string) (num int64, posts []model.Post, err error) {
	num, err = db.FilterTable(filters, &posts, DBTableNamePost)
	if err == nil {
		for k, _ := range posts {
			db.RelodeComment(&posts[k])
		}
	}
	return
}

func (db *DB) GetPost(postID string) (post model.Post, err error) {

	return
}

func (db *DB) RelodeComment(post *model.Post) {
	err := db.ORM.Model(post).Related(&post.Comments).Error
	if err != nil {
		glog.Errorf("Post reload related error: %s.", err.Error())
	}
}
