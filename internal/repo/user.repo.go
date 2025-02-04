package repo

import (
	"myproject/global"
	"myproject/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func (u *userRepository) GetUserByEmail(email string) bool {
	// var user model.GoDbUser
	//
	// result := global.Mdb.
	// 	Table(TableNameGoCrmUser).
	// 	Where("email = ?", email).
	// 	First(&user)
	//
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return false
	// }
	//
	// if result.Error != nil {
	// 	return false
	// }
	//
	// return true
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
