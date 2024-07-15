package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string
	Slug  string
}

// Tablo oluşmasını sağlıyoruz.
func (category Category) Migrate() {
	OpenDB()

	DB.AutoMigrate(&category)
}

// Ekleme ile ilgili method
func (category Category) Add() {
	OpenDB()

	DB.Create(&category)
}

// Veri çekme işlemi için
// where: hangi verinin geleceği belli olmadığı için garantiye alıp interface yaptık
func (category Category) Get(where ...interface{}) Category {
	OpenDB()

	DB.First(&category, where...)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	OpenDB()

	var categories []Category
	DB.Find(&categories, where...)
	return categories
}

func (category Category) Update(column string, value ...interface{}) {
	OpenDB()
	DB.Model(&category).Update(column, value)
}

func (category Category) Updates(data Category) {
	OpenDB()
	DB.Model(&category).Updates(data)
}

func (category Category) Delete() {
	OpenDB()

	DB.Delete(&category, category.ID)
}
