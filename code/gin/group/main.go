package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger1() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("logger 1 before request")

		c.Next()

		log.Println("logger 1 after request")
	}
}

func Logger2() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("logger 2 before request")

		c.Next()

		log.Println("logger 2 after request")
	}
}

func registerAPI(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(Logger2())

	api.GET("/test", func(c *gin.Context) {
		log.Println("trigger logger1 and logger2")
	})
}

// 1. test URL http://localhost:8080/api
// 2. test URL http://localhost:8080/api/test
func main() {
	r := gin.New()

	r.Use(Logger1())
	
	registerAPI(r)

	r.Run(":8080")
}
