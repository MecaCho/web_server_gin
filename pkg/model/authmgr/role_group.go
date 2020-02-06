package authmgr

// RoleGroup role group data model
type RoleGroup struct {
	ID          int64   `json:"id" orm:"auto;pk;unique;column(id)"`
	Name        string  `json:"name" orm:"size(64);unique;"`
	Description string  `json:"description" orm:"size(256);null"`
	GroupType   string  `json:"group_type"`
	Roles       []*Role `json:"roles" orm:"reverse(many)"`
	Users       []*User `json:"users" orm:"reverse(many)"`
}

// TableName 数据表名称
func (rg *RoleGroup) TableName() string {
	return DBTableNameRoleGroup
}
