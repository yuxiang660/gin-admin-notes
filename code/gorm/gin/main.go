package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Model gorm.Model

// Product ...
type Product struct {
	Model
	Code string `json:"code"`
	Price uint `json:"price"`
}

// ProductModel ...
type ProductModel struct {
	db *gorm.DB
}

// AutoMigrate ...
func (m *ProductModel) AutoMigrate(values ...interface{}) error {
	return m.db.AutoMigrate(values...).Error
}

// Create creates a user with username and password.
func (m *ProductModel) Create(c *gin.Context) {
	var p Product
	if err := c.ShouldBind(&p); err != nil {
		c.String(400, "error")
		return
	}

	fmt.Println("Code:", p.Code)
	fmt.Println("Price:", p.Price)

	m.db.Create(&p)

	var product Product
	m.db.First(&product, 1)
	fmt.Println("id 1:", product)

	m.db.First(&product, "code = ?", "L1212")
	fmt.Println("code L1212:", product)

	m.db.Delete(&product)

	c.String(200, "success")
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db = db.Debug()

	model := ProductModel{db: db}

	err = model.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to auto migrate")
	}

	app := gin.Default()
	app.POST("/create", model.Create)

	app.Run(":12345")
}
