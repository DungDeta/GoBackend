package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"go.uber.org/zap"
	"myproject/global"
)

func InitMysqlc() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.DbName)
	db, err := sql.Open("mysql", s)
	if err != nil {
		global.Logger.Error("Mysql connect error", zap.Error(err))
		panic(err) // Tắt sv luôn nếu không kết nối được
	}
	global.Logger.Info("Mysql connect success")
	global.Mdbc = db
	// SetPool() vì tính chất của TCP là khi kết nối xong sẽ đóng lại nên set pool sẽ duy trì liên tục 1 pool kết nối
	// SetPool()
	// migrateTables()
}

func SetPoolc() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("Mysql get db error", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)                      // Số lượng kết nối tối đa được giữa 2 lần sử dụng
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)                      // Số lượng kết nối tối đa được mở
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) // Thời gian tối đa mà một kết nối có thể được sử dụng
}

func migrateTablesc() {
	// // err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	// if err != nil {
	// 	global.Logger.Error("Migrate tables error", zap.Error(err))
	// 	panic(err)
	// }
}
