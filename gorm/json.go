package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// AddressDetail 自定义json类型
type AddressDetail []string

func (a AddressDetail) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(a)
	return string(bytes), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
// 这里的 *AddressDetail 必须使用指针
func (a *AddressDetail) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	address := AddressDetail{}
	err := json.Unmarshal(bytes, &address)
	*a = address
	return err
}

// BizCustomer 业务客户
// https://github.com/go-gorm/datatypes
// https://gorm.io/zh_CN/docs/data_types.html
type BizCustomer struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	// 表示 json 必须要使用自定义类型，mysql5.x 不支持 json 类型
	// 如果强制修改 datatypes.JSON 中的 GormDBDataType 方法，会导致 json 查询失效，因为 json 对应的函数不存在
	// 比如在 mysql5.7 将 GormDBDataType 修改为 longtext，datatypes.JSONQuery 查询失效
	Address       datatypes.JSON `json:"address" gorm:"column:address"`
	Detail        datatypes.JSON `json:"detail" gorm:"column:detail"`
	AddressDetail AddressDetail  `json:"addressDetail" gorm:"column:address_detail;type:json"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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
		AddressDetail: []string{"Moscow", "Krasnodar"},
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
	db.Where("id = ?", customer.ID).First(&customerFromDb3, datatypes.JSONQuery("detail").HasKey("phone"))
	println(ToJsonWithIndent(customerFromDb3))

	// 查询不到
	var customerFromDb4 BizCustomer
	db.Where("id = ?", customer.ID).First(&customerFromDb4, datatypes.JSONQuery("detail").HasKey("phone1"))
	println(ToJsonWithIndent(customerFromDb4))
}
