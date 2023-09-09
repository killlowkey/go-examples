package main

import (
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// AfterFind 在查询之后调用
func (p *Product) AfterFind(*gorm.DB) error {
	log.Println("AfterFind: ", ToJson(p))
	return nil
}

// BeforeCreate 再被创建之前调用，此时 p 里面的值并未填充，为字段默认值
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate: ", ToJson(p))
	return nil
}

// AfterCreate 在创建之后调用,但是事务并未提交，返回错误，则事务回滚
func (p *Product) AfterCreate(tx *gorm.DB) error {
	log.Println("AfterCreate: ", ToJson(p))
	return nil
}

// BeforeUpdate 在更新之前调用，此时 p 里面的值并未填充，为字段默认值
func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	log.Println("BeforeUpdate: ", ToJson(p))
	return nil
}

// AfterUpdate 在更新之后调用,但是事务并未提交，返回错误，则事务回滚
func (p *Product) AfterUpdate(tx *gorm.DB) error {
	log.Println("AfterUpdate: ", ToJson(p))
	return nil
}

// BeforeDelete 在删除之前调用，此时 p 里面的值并未填充，为字段默认值
func (p *Product) BeforeDelete(tx *gorm.DB) error {
	log.Println("BeforeDelete: ", ToJson(p))
	return nil
}

// AfterDelete 在删除之后调用,但是事务并未提交，返回错误，则事务回滚
func (p *Product) AfterDelete(tx *gorm.DB) error {
	log.Println("AfterDelete: ", ToJson(p))
	return nil
}

func hookExample() {
	//db.Callback().Create().Before("gorm:create").Register("before_create", beforeCreate)
	//db.Callback().Create().After("gorm:create").Register("after_create", afterCreate)

	// Migrate the schema
	err := db.AutoMigrate(&Product{})
	if err != nil {
		log.Panicln("migrate product schema error: ", err)
	}

	// create: 会触发 create hook: BeforeCreate, AfterCreate
	err = db.Model(&Product{}).Create(&Product{Code: UUID(), Price: 100}).Error
	if err != nil {
		log.Panicln("hookExample create product error: ", err)
	}

	//  find: 会触发 find hook: AfterFind
	var products []Product
	err = db.Model(&Product{}).Limit(5).Find(&products).Error
	if err != nil {
		log.Panicln("hookExample find product error: ", err)
	}
	log.Println("hookExample find product: ", ToJson(products))

	//  update: 会触发 update hook: BeforeUpdate, AfterUpdate
	err = db.Model(&Product{}).Where("id = ?", 1).Updates(Product{Price: 200}).Error
	if err != nil {
		log.Panicln("hookExample update product error: ", err)
	}

	// delete: 会触发 delete hook: BeforeDelete, AfterDelete
	err = db.Model(&Product{}).Where("id = ?", 1).Delete(&Product{}).Error
	if err != nil {
		log.Panicln("hookExample delete product error: ", err)
	}
}

// crudExample
// https://gorm.io/docs/connecting_to_the_database.html#SQLite
func crudExample() {
	// Migrate the schema
	err := db.AutoMigrate(&Product{})
	if err != nil {
		log.Panicln("migrate product schema error: ", err)
	}

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
