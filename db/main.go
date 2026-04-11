package main

import (
	"log"

	"leetcode-go/db/mysql"
)

func main() {
	user, password, dbName := "root", "Jqqi,.88", "test_db"

	// 1. 连接 MySQL（不指定数据库）
	db, err := mysql.ConnectWithoutDB(user, password)
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}

	// 2. 创建数据库
	if err := mysql.CreateDatabase(db, dbName); err != nil {
		log.Fatalf("创建数据库失败: %v", err)
	}
	log.Printf("数据库 %s 创建成功\n", dbName)

	// 3. 关闭连接，重新连接指定数据库
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db, err = mysql.ConnectLocalMySQL(user, password, dbName)
	if err != nil {
		log.Fatalf("连接数据库 %s 失败: %v", dbName, err)
	}

	// 4. 创建 user 表
	if err := mysql.CreateUserTable(db); err != nil {
		log.Fatalf("创建 user 表失败: %v", err)
	}

	log.Println("user 表创建成功")
}
