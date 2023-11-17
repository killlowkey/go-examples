package main

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// BizCustomer 业务客户
// https://github.com/go-gorm/datatypes
// https://gorm.io/zh_CN/docs/data_types.html
type BizCustomer struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	// 表示 json 必须要使用自定义类型，mysql5.x 不支持 json 类型
	// 如果强制修改 datatypes.JSON 中的 GormDBDataType 方法，会导致 json 查询失效，因为 json 对应的函数不存在
	// 比如在 mysql5.7 将 GormDBDataType 修改为 longtext，datatypes.JSONQuery 查询失效
	Address   datatypes.JSON `json:"address" gorm:"column:address"`
	Detail    datatypes.JSON `json:"detail" gorm:"column:detail`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b BizCustomer) TableName() string {
	return "biz_customer"
}

func marshal(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return bytes
}

func jsonExample() {
	err := db.AutoMigrate(&BizCustomer{})
	if err != nil {
		panic(err)
	}

	address := []string{"Moscow", "Krasnodar"}

	customer := BizCustomer{
		Name:    "Gleb",
		Address: datatypes.JSON(marshal(address)),
		Detail: datatypes.JSON(marshal(map[string]string{
			"phone": "123456789",
			"email": "2860072080@gmail.com",
		})),
	}
	db.Create(&customer)

	var customerFromDb BizCustomer
	db.First(&customerFromDb, customer.ID)
	println(ToJsonWithIndent(customerFromDb))

	//address = append(address, "Sochi")
	//customer.Address = marshal(address)
	//db.Save(&customer)

	//var customerFromDb2 BizCustomer
	//db.First(&customerFromDb2, customer.ID)
	//println(ToJsonWithIndent(customerFromDb2))

	var customerFromDb3 BizCustomer
	db.First(&customerFromDb3, datatypes.JSONQuery("detail").HasKey("phone"))
	println(ToJsonWithIndent(customerFromDb3))

	var customerFromDb4 BizCustomer
	db.First(&customerFromDb4, datatypes.JSONQuery("detail").HasKey("phone1"))
	println(ToJsonWithIndent(customerFromDb4))
}
