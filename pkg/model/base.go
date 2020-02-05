package model

import (
	"time"
)

// BaseCommonModel base meta data
type BaseCommonModel struct {
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;unique;Column:id"`
	Name        string    `json:"name" gorm:"size:64" binding:"lt=64,nameReg"`
	Description string    `json:"description" gorm:"size:256;null" binding:"lt=256"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null" time_format:"2006-01-02"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"null" time_format:"2006-01-02"`
	Creator     string    `json:"creator" gorm:"size:32" binding:"lt=32"`
	Modifier    string    `json:"modifier" gorm:"size:32;null" binding:"lt=32"`
}

// type User struct {
// 	gorm.Model
// 	Name         string
// 	Age          sql.NullInt64
// 	Birthday     *time.Time
// 	Email        string  `gorm:"type:varchar(100);unique_index"`
// 	Role         string  `gorm:"size:255"` // set field size to 255
// 	MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
// 	Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
// 	Address      string  `gorm:"index:addr"` // create index with name `addr` for address
// 	IgnoreMe     int     `gorm:"-"` // ignore this field
// }
