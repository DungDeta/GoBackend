package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Kết nối database
	dsn := "root:root1234@tcp(localhost:33306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn))

	// Tạo generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/model/query", // Thư mục đầu ra
		ModelPkgPath: "./internal/model",       // Thư mục model
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Sử dụng database
	g.UseDB(db)

	// Generate model cho các bảng
	g.GenerateModel("go_db_users") // Tạo model cho bảng users
	g.GenerateModel("orders")      // Tạo model cho bảng orders

	// Thực thi code generation
	g.Execute()
}
