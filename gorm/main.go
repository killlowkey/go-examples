package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// initSQLite
// 1. 纯GO驱动：go get -u github.com/glebarez/sqlite
// 2. CGO 驱动：go get -u gorm.io/driver/sqlite
func initSQLite() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
}

// initMySQL go get -u gorm.io/driver/mysql
func initMySQL() {
	var err error
	dsn := "root:root@tcp(172.16.60.77:3307)/gva?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
}

func init() {
	//initSQLite()
	initMySQL()
}

func main() {
	//hookExample()
	//sessionExample()
	//hookCustomExample()
	//crudExample()
	//hasMany()
	jsonExample()
}
