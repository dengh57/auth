package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:1234@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 300,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接成功")
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
