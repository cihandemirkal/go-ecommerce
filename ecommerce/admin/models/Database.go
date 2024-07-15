package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dns string = "root:root@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func OpenDB() {

	var err error
	DB, err = gorm.Open(mysql.Open(Dns), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

}
