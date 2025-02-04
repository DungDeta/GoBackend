package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id       int64  `gorm:"column:id; not null; primaryKey; autoIncrement; comment:id is primary key"`
	Rolename string `gorm:"column:rolename; type:varchar(255); not null; index:idx_username"`
	RoleNote string `gorm:"column:role_note; type:text; not null"`
}

func (r Role) TableName() string {
	return "go_db_roles"
}
