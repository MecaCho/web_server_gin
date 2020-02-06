package authmgr

// UserGroupRole user's  role in a group
type UserGroupRole struct {
	ID      int64   `json:"id" orm:"auto;pk;unique;column(id)"`
	Users   []*User `json:"users" orm:"reverse(many)"`
	GroupID int64   `json:"group_id"`
	RoleID  int64   `json:"role_id"`
}

// TableName 数据表名称
func (r *UserGroupRole) TableName() string {
	return DBTableNameUserGroupRole
}
