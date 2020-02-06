package dao

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
	"web_server_gin/pkg/model"
)

func TestDB_FilterTable(t *testing.T) {
	// filters := map[string][]string{}
	// // filters[]
	// if filters != nil {
	// 	fmt.Println(filters, "not nil")
	// } else {
	// 	fmt.Println(filters, "nil")
	// }

	dbCon := "root:xxxxxxx@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	db, err := InitDB(dbCon)
	if err != nil {
		t.Error(err)
	}
	// var modelList model.Post
	// filterMap := make(map[string]interface{})
	// filterMap["id"] = 5
	// err = db.ORM.Limit(10).
	// 	Where(filterMap).
	// 	Offset(0).
	// 	Order("created_at").
	// 	Find(&modelList).
	// 	Error
	// fmt.Println("result: ",err, modelList)

	var post []model.Post
	filters := make(map[string][]string)
	filters["id"] = []string{"5"}
	num, err := db.FilterTable(filters, &post, DBTableNamePost)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(post, err, num)
}
