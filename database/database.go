package database

import (
	"IrisBlog/model"
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const db_type int = 1

func sqlite3() *gorm.DB {

	db, err := gorm.Open("sqlite3", "/tmp/IrisBlog.db")

	if err != nil {
		fmt.Println(err)
		panic("无法连接数据库")
	}
	fmt.Println("连接sqlite3数据库成功")

	return db

}

func mysql() *gorm.DB {

	db, err := gorm.Open("mysql", "root:root@(localhost)/irisblog?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		panic("无法连接数据库")
	}
	fmt.Println("连接mysql数据库成功")

	return db

}

func Db() *gorm.DB {

	switch db_type {
	case 0:
		db := mysql()
		//单数模式
		db.SingularTable(true)
		// 创建默认表
		db.AutoMigrate(&model.User{})
		return db
	case 1:
		db := sqlite3()
		//单数模式
		db.SingularTable(true)
		// 创建默认表
		db.AutoMigrate(&model.User{})
		return db
	default:
		panic("未知的数据库")
	}

	// 逻辑结束后关闭数据库
	// defer func() {
	// 	_ = db.Close()
	// }()

}
