package main

import (
	"fmt"
	"gorm.io/gorm"
)

func hasOneExample() {
	type Company struct {
		gorm.Model
		Name   string
		UserID uint
		//UserName string
	}

	type User struct {
		gorm.Model
		Name    string
		Company Company `gorm:"foreignKey:UserID"` // User=>id 关联到 Company=>UserId
		//Company Company `gorm:"foreignKey:UserName"` // User=>id 关联到 Company=>UserName
	}

	err := db.AutoMigrate(&User{}, &Company{})
	if err != nil {
		panic(err)
	}

	// 创建
	user := User{
		Name: "ray",
		Company: Company{
			Name: "无穷大科技有限公司",
		},
	}
	db.Create(&user)

	// 查询
	var res User
	// Where("id = ?", 1)
	err = db.Model(&User{}).Preload("Company").First(&res, 1).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(ToJsonWithIndent(&res))
}
