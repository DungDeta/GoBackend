package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid      uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; primaryKey; index:idx_uuid, unique"`
	Username  string    `gorm:"column:username; type:varchar(255); not null; index:idx_username"`
	IsActived bool      `gorm:"column:is_actived; type:boolean; not null; default:false"`
	Role      []Role    `gorm:"many2many:go_user_roles;"`
}

func (u User) TableName() string {
	return "go_db_users"
}
