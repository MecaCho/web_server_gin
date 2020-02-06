package authmgr

import "time"

// Role db model
type Role struct {
	ID          int64        `json:"id" orm:"auto;pk;unique;column(id)"`
	Name        string       `json:"name" orm:"size(64);unique;"`
	Description string       `json:"description" orm:"size(256);null"`
	Users       []*User      `json:"users" orm:"reverse(many)"`
	RoleGroups  []*RoleGroup `json:"role_groups" orm:"rel(m2m)"`
	CreatedAt   time.Time    `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty" orm:"auto_now;type(datetime);null;default(null)"`
	Creator     string       `json:"creator" orm:"size(32);"`
	Modifier    string       `json:"modifier" orm:"size(32);null"`
}

// TableName 数据表名称
func (r *Role) TableName() string {
	return DBTableNameRole
}
