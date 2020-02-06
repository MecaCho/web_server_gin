package authmgr

// UserGroup user group data model
type UserGroup struct {
	ID          int64   `json:"id" orm:"auto;pk;unique;column(id)"`
	Name        string  `json:"name" orm:"size(64);unique;"`
	Description string  `json:"description" orm:"size(256);null"`
	GroupType   string  `json:"group_type"`
	ParentGroup int64   `json:"parent_group"`
	Users       []*User `json:"users" orm:"reverse(many)"`
}

// TableName 数据表名称
func (ug *UserGroup) TableName() string {
	return DBTableNameUserGroup
}

// TableUnique 多字段唯一键
func (ug *UserGroup) TableUnique() [][]string {
	return [][]string{
		[]string{"name", "group_type"},
	}
}
