package dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"strconv"
)

// // InitDB init db connection
// func InitDB(conn string) (*sql.DB, error) {
// 	db, err := sql.Open("postgres", conn)
//
// 	if err != nil {
// 		glog.Errorf("Init DB error, %s", err.Error())
// 	}
// 	return db, err
// }

// DB db handle
type DB struct {
	ORM       gorm.DB
	DebugMode bool
}

//
// // RegisterDataBase register database
// func RegisterDataBase(dbConn, aliasName string) (err error) {
// 	dbUsr, dbPwd, dbEndpoint, err := util.ParserConStr(dbConn)
// 	if err != nil {
// 		glog.Errorf("Decode db auth error, %s", err.Error())
// 		return
// 	}
// 	dbConnDataSource := fmt.Sprintf("%s:%s@%s", dbUsr, dbPwd, dbEndpoint)
// 	err = orm.RegisterDataBase(aliasName, "mysql",
// 		dbConnDataSource, 30, 30)
// 	if err != nil {
// 		glog.Errorf("Register data base error : %s.", err.Error())
// 		return
// 	}
// 	glog.Infof("register data base :%s.", aliasName)
// 	return
// }
//
// // InitDBConnect init db connect
// func InitDBConnect(dbConn, hostDBconn string, isDebugMode bool) {
// 	err := RegisterDataBase(dbConn, meta.DBAliasDefault)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	err = RegisterDataBase(hostDBconn, meta.DBAliasLowLoad)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = RegisterDataBase(strings.Replace(hostDBconn, meta.DBAliasLowLoad, meta.DBAliasOps, 1), meta.DBAliasOps)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	err = RegisterDataBase(strings.Replace(hostDBconn, meta.DBAliasLowLoad, meta.DBAliasCC, 1), meta.DBAliasCC)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	RegisterModel()
// 	orm.RunSyncdb("default", false, true)
// 	orm.Debug = isDebugMode
// 	ServerDB.DebugMode = isDebugMode
// 	ServerDB.ORM = orm.NewOrm()
//
// 	glog.Infof("Init db connect, %s.", dbConn)
// }
//

//
// // ListColumnNames list all columns name of table
// func (db *DB) ListColumnNames(tableName string) (columns []string) {
//
// 	// TODO query table columns
// 	// ...err := db.ORM.Raw("select COLUMN_NAME from information_schema.COLUMNS where table_name ="application";")
// 	return
// }
//
// // ConvertFilterValue convert filter values
// func ConvertFilterValue(value string) interface{} {
//
// 	boolRet, err := strconv.ParseBool(value)
// 	if err == nil {
// 		return boolRet
// 	}
// 	IntRet, err := strconv.ParseInt(value, 10, 64)
// 	if err == nil {
// 		return IntRet
// 	}
// 	return value
// }
//
// // DeleteResource delete a resource
// func (db *DB) DeleteResource(resourceID int64, resourceDB interface{}, tableName string) (err error) {
// 	var id int64
//
// 	defer func() {
// 		if err != nil {
// 			glog.Infof("Delete resource %d error: %+v.", resourceID, err)
// 		} else {
// 			glog.Infof("Delete resource (%d) success: %d.", resourceID, id)
// 		}
// 	}()
// 	id, err = db.ORM.QueryTable(resourceDB).Filter(DbID, resourceID).Delete()
// 	return err
// }
//
// CreateResource create a resource
func (db *DB) CreateResource(resourceDB interface{}) (err error) {
	defer func() {
		if err != nil {
			glog.Infof("Create resource error: %+v.", err)
		} else {
			glog.Infof("Create resource (%+v) success.", resourceDB)
		}
	}()
	err = db.ORM.Create(resourceDB).Error
	return
}

// UpdateResource update resource in db
func (db *DB) UpdateResource(resourceDB interface{}, columns ...string) (err error) {
	var id int64
	defer func() {
		if err != nil {
			glog.Infof("Update resource error: %+v, detail: %+v, id: %d.", err, resourceDB, id)
		} else {
			glog.Infof("Update resource (%d) success, detail: %+v.", id, columns)
		}
	}()

	err = db.ORM.Update(resourceDB).Error
	return
}

// FilterTable filter resource from db
func (db *DB) FilterTable(filters map[string][]string, modelList interface{}, tableName string) (num int64, err error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("List resources in db error, %+v.", err)
		}
	}()

	limit, offset, orderValue := int64(1000), int64(0), "created_at"
	num, err = int64(0), nil

	// typeSlice := reflect.TypeOf(modelList).Elem()

	// if typeSlice.Kind() != reflect.Slice {
	// 	return 0, fmt.Errorf("query db table model type error, %s", typeSlice.Kind())
	// }

	// modelInterface := reflect.ValueOf(modelList).Interface()
	// jsonTagMap := GetStructJSONTag(modelInterface)

	qs := db.ORM

	for k, v := range filters {
		if len(v) < 1 {
			continue
		}
		value := v[0]
		if k == "limit" {
			limit, _ = strconv.ParseInt(value, 10, 64)
			glog.Infof("Filter table really limit : %d.", limit)
			delete(filters, k)
			continue
		} else if k == "offset" {
			offset, _ = strconv.ParseInt(value, 10, 64)
			delete(filters, k)
			continue
		} else if k == "sorted" {
			orderValue = value
			delete(filters, k)
			continue
		}

	}

	qs.Find(modelList).Count(&num)

	if len(filters) != 0 {
		err = qs.
			Limit(limit).
			Where(filters).
			Offset(offset).
			Order(orderValue).
			Find(modelList).
			Error
	} else {
		err = qs.
			Limit(limit).
			Offset(offset).
			Order(orderValue).
			Find(modelList).
			Error
	}

	// num = int64(len(modelList))
	// totalCount := num
	// if limit != 1000 || offset != 0 {
	// 	qs.Count(&totalCount)
	// }
	totalCount := num

	glog.Infof("List %s list,%d,%d,%s,filters: %+v, num: %d,count: %d, %+v.",
		tableName, limit, offset, orderValue, filters, num, totalCount, modelList)

	if err != nil {
		return 0, fmt.Errorf("query db error, please check parameters, detail: %s", err.Error())
	}
	return num, err
}
