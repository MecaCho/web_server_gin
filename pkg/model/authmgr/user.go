package authmgr

import "time"

// User db model
type User struct {
	ID               int64            `json:"id" orm:"auto;pk;unique;column(id)"`
	Name             string           `json:"name" orm:"size(64);unique;"`
	Description      string           `json:"description" orm:"size(256);null"`
	Email            string           `json:"email" orm:"size(128);unique;"`
	PhoneNumber      string           `json:"phone_number" orm:"size(16);"`
	RealName         string           `json:"real_name" orm:"size(64);"`
	TokenExpired     time.Time        `json:"token_expired" orm:"null"`
	PasswordExpired  time.Time        `json:"password_expired" orm:"auto_now;type(datetime);null;default(null)"`
	CreatedAt        time.Time        `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt        time.Time        `json:"updated_at,omitempty" orm:"auto_now;type(datetime);null;default(null)"`
	Creator          string           `json:"creator" orm:"size(32);"`
	Modifier         string           `json:"modifier" orm:"size(32);null"`
	VerifyCode       string           `json:"verify_code" orm:"size(8);null"`
	VerifyCodeExpire time.Time        `json:"verify_code_expire" orm:"auto_now;type(datetime);null;default(null)"`
	UserType         string           `json:"user_type" orm:"size(32);"`
	Roles            []*Role          `json:"roles" orm:"rel(m2m)"`
	Groups           []*UserGroup     `json:"groups" orm:"rel(m2m)"`
	RoleGroups       []*RoleGroup     `json:"role_groups" orm:"rel(m2m)"`
	GroupRoles       []*UserGroupRole `json:"group_roles" orm:"rel(m2m)"`
}

// TableName 数据表名称
func (u *User) TableName() string {
	return DBTableNameUser
}
