package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Course struct {
	CourseId   uint `gorm:"primaryKey"` // 配置主键
	CourseName string
	Students   []Student `gorm:"many2many:student_course;"` // 这里的many2many是表名，与 student 一样
}

type Student struct {
	gorm.Model
	Name    string
	Age     int
	Courses []Course `gorm:"many2many:student_course;"` // 与 course 一样。插入时，只需要携带课程id
}

func many2manyExample() {
	err := db.AutoMigrate(&Student{}, &Course{})
	if err != nil {
		panic(err)
	}

	db.Model(&Course{}).Create(&Course{CourseName: "Java"})
	db.Model(&Course{}).Create(&Course{CourseName: "Go"})
	db.Model(&Course{}).Create(&Course{CourseName: "Python"})

	var courses []Course
	courses = append(courses, Course{CourseId: 1})
	courses = append(courses, Course{CourseId: 2})
	courses = append(courses, Course{CourseId: 3})

	db.Model(&Student{}).Create(&Student{Name: "John", Age: 18, Courses: courses[0:1]})
	db.Model(&Student{}).Create(&Student{Name: "Ray", Age: 20, Courses: courses[1:]})

	// many2many 是主键与主键的关联，只需要指定好对应 struct 的主键和配置 many2many 即可
	var res Student
	// 严格按照这种顺序写
	err = db.Preload("Courses").Model(&Student{}).Where("id = ?", 2).First(&res).Error
	if err != nil {
		panic(err)
	}
	println(ToJsonWithIndent(&res))

	res.Courses = append(res.Courses, Course{CourseId: 2})
	err = db.Preload("Courses").Model(&Student{}).Updates(&res).Error
	if err != nil {
		panic(err)
	}

	var course Course
	err = db.Preload("Students").Model(&Course{}).Where("course_id = ?", 2).First(&course).Error
	if err != nil {
		panic(err)
	}
	println(ToJsonWithIndent(&course))
}

func many2manyExample2() {
	type Language struct {
		gorm.Model
		Name string
	}

	// User 拥有并属于多种 language，`user_languages` 是连接表
	type IUser struct {
		gorm.Model
		Languages []Language `gorm:"many2many:user_languages;"`
	}

	err := db.AutoMigrate(&IUser{}, &Language{})
	if err != nil {
		panic(err)
	}

	db.Create(&IUser{
		Languages: []Language{
			{Name: "Go"},
			{Name: "Java"},
			{Name: "Python"},
		},
	})

	var res IUser
	db.Model(&IUser{}).Preload("Languages").First(&res)
	fmt.Println(ToJsonWithIndent(&res))
}
