package main

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// User 用户实体
// https://gorm.io/zh_CN/docs/has_many.html
type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Age       string
	Address   string
	Book      []Book
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

// Book 书籍实体
type Book struct {
	ID        uint `gorm:"primarykey"`
	UserId    int  // 必须要有这个，才能进行关联查询
	Name      string
	Perice    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Book) TableName() string {
	return "book"
}

func autoMigrateTable() {
	err := db.AutoMigrate(&User{}, &Book{})
	if err != nil {
		panic(err)
	}
}

func createUser() {
	db.Create(&User{Name: "jinzhu", Age: "18", Address: "Shanghai", Book: []Book{
		{
			Name:   "Go in Action",
			Perice: 100,
		},
	}})
}

func findUserById(id int) (*User, error) {
	var user User
	// 首先要先写 Model(&User{})
	// Preload("Book") => Book 是 User 中的 Book 字段名，必须要这样填，才能查询数据
	err := db.Model(&User{}).Preload("Book").Where("users.id = ?", id).First(&user).Error
	//err := db.Model(&User{}).Preload("Book").First(&user, id).Error
	return &user, err
}

func saveUser(user *User) error {
	// 追加书籍，会更新
	user.Book = append(user.Book, Book{
		Name:   "Java in action",
		Perice: 50,
	})
	return db.Save(&user).Error
}

func hasMany() {
	autoMigrateTable()

	user, _ := findUserById(1)
	log.Println(ToJsonWithIndent(user))
	//
	//if len(user.Book) >= 2 {
	//	// 删除第二本书籍
	//	db.Delete(&Book{}, "id = ?", user.Book[1].ID)
	//}
	//
	//user, _ = findUserById(1)
	//log.Println(ToJsonWithIndent(user))

	//row := db.Model(&User{}).Joins("JOIN book on users.id = book.user_id").Select("book.name, book.user_id").Limit(1).Row()
	//var name string
	//var age int
	//row.Scan(&name, &age)
	//log.Printf("name=%s userId=%d\n", name, age)

	//rows, _ := db.Model(&User{}).Joins("JOIN book on users.id = book.user_id").Select("book.name, book.user_id").Rows()
	//var name string
	//var age int
	//rows.Scan()
	//log.Printf("name=%s userId=%d\n", name, age)

	//var res map[string]any
	row := db.Exec("select * from users where id = ? limit 1", 1).Row()
	row.Scan()
	//log.Printf(ToJsonWithIndent(&res))
}

type Res struct {
	Name   string
	UserId int
}

func (r *Res) Scan(src any) error {
	log.Println(src)
	return nil
}
