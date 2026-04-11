package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewMySQLConnection(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 连接池设置
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func ConnectLocalMySQL(user, password, dbName string) (*gorm.DB, error) {
	return NewMySQLConnection(Config{
		Host:     "localhost",
		Port:     3306,
		User:     user,
		Password: password,
		DBName:   dbName,
	})
}

// ConnectWithoutDB 连接 MySQL 但不指定数据库
func ConnectWithoutDB(user, password string) (*gorm.DB, error) {
	return NewMySQLConnection(Config{
		Host:     "localhost",
		Port:     3306,
		User:     user,
		Password: password,
		DBName:   "",
	})
}

// CreateDatabase 创建数据库（如果不存在）
func CreateDatabase(db *gorm.DB, dbName string) error {
	return db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName)).Error
}
