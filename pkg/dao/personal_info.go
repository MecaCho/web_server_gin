package dao

import (
	"web_server_gin/pkg/model"
	"fmt"
	"github.com/golang/glog"
)

func (db *DB) GetPersonalInfo(filters map[string][]string) (num int64, info model.PersonaInfo, err error) {
	var infos []model.PersonaInfo
	num, err = db.FilterTable(filters, &infos, model.DBTableNamePersonalInfo)
	glog.Infof("List personal infos: %+v.", infos)
	if num == 0 {
		err = fmt.Errorf("personal info not found")
		return
	}
	if err == nil {
		info = infos[0]
	}
	return
}
