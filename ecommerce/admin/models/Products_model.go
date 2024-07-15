package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID  int
	Name        string
	Description string
	Price       int
	Stock       int
	Picture_url string
}

// Tablo oluşmasını sağlıyoruz.
func (product Product) Migrate() {
	OpenDB()

	DB.AutoMigrate(&product)
}

// Ekleme ile ilgili method
func (product Product) Add() {
	OpenDB()

	DB.Create(&product)
}

// Veri çekme işlemi için
// where: hangi verinin geleceği belli olmadığı için garantiye alıp interface yaptık
func (product Product) Get(where ...interface{}) Product {
	OpenDB()

	DB.First(&product, where...)
	return product
}

func (product Product) GetAll(where ...interface{}) []Product {
	OpenDB()

	var products []Product
	DB.Find(&products, where...)
	return products
}

func (product Product) Update(column string, value ...interface{}) {
	OpenDB()
	DB.Model(&product).Update(column, value)
}

func (product Product) Updates(data Product) {
	OpenDB()
	DB.Model(&product).Updates(data)
}

func (product Product) Delete() {
	OpenDB()

	DB.Delete(&product, product.ID)
}
