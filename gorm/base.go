package main

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// AfterFind 在查询之后调用
func (p *Product) AfterFind(tx *gorm.DB) error {
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

// hookExample https://gorm.io/zh_CN/docs/hooks.html
func hookExample() {
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

func hookCustomExample() {
	// RegisterDefaultCallbacks
	err := db.Callback().Create().Before("gorm:create").Register("custom_before_create", func(tx *gorm.DB) {
		log.Println("custom_before_create: ")
	})
	if err != nil {
		log.Panicln("hookCustomExample create error: ", err)
	}
	//db.Callback().Create().After("gorm:create").Register("after_create", afterCreate)
	db.Create(&Product{Code: UUID(), Price: 100})
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

// txAutoExample 闭包自动提交事务，返回 error 则回滚，返回 nil 则提交
func txAutoExample() {
	err := db.Transaction(func(tx *gorm.DB) error {
		// 开启事务
		err := tx.Create(Product{Code: "T42", Price: 200}).Error
		if err != nil {
			return errors.New("product's price is too much")
		}

		return nil
	})

	if err != nil {
		log.Println("transaction error: ", err)
	}
}

// txManualExample 手动提交事务
func txManualExample() {
	// 开始事务
	tx := db.Begin()
	err := tx.Create(&Product{Code: "T42", Price: 200}).Error

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func sessionExample() {
	// DryRun 调试使用，不会清除 SQL 语句
	session := db.Session(&gorm.Session{DryRun: true})
	session.Create(&Product{Code: "D42", Price: 100})
}

func resultExample() {
	// 返回的 db 禁止复用，会导致 sql 生成有问题
	tx := db.Create(&Product{Code: "D42", Price: 100})

	// 该 sql 影响行数
	rowsAffected := tx.RowsAffected
	log.Println("rows affected: ", rowsAffected)

	// 执行错误，推荐这种方式处理返回错误，record 未发现，返回 gorm.RecordNotFoundError
	err := tx.Error
	if err != nil {
		log.Println(err)
	}

	// 返回原生记录 record
	rows, err := tx.Rows()
	log.Println("rows: ", rows)
}

func clauseExample() {
	var (
		products []Product
		limit    = 10
	)
	// db.Offset(1).Limit(limit).Find(&products)
	db.Clauses(clause.Limit{Offset: 1, Limit: &limit}).Find(&products)
}
