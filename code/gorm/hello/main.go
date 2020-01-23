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
	db.Create(&Product{Code: "L2222", Price: 2222})

	var p1 Product
	db.First(&p1, "code = ?", "L1212")
	fmt.Println("code L1212:", p1)

	var p2 Product
	db.First(&p2, "code = ?", "L2222")
	fmt.Println("code L2222:", p2)

	var p3 Product
	db.First(&p3, "code = ?", "L3333")
	fmt.Println("code L3333:", p3)

	err = db.Delete(&p1).Error
	if err != nil {
		fmt.Println("d1 err")
		return
	}
	err = db.Delete(&p2).Error
	if err != nil {
		fmt.Println("d2 err")
		return
	}
	err = db.Delete(&p3).Error
	if err != nil {
		fmt.Println("d3 err")
		return
	}
}