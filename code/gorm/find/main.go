package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&Product{Code: "L1212", Price: 2222})
	db.Create(&Product{Code: "L1212", Price: 3333})

	var products []Product

	db.Find(&products, "code = ?", "L1212")
	fmt.Println(products)

	err = db.Delete(&products).Error
	if err != nil {
		fmt.Println("delete!")
		return
	}
}