package mysql

import "gorm.io/gorm"

// User 模型
type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(100);not null"`
	Age  int    `gorm:"default:0"`
}

// CreateUserTable 创建 user 表
func CreateUserTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
