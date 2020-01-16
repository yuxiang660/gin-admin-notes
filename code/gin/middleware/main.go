package main

import (
	"log"
	"time"
	
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		log.Println("Before request: set example to 12345")

		c.Next()

		latency := time.Since(t)
		log.Println("After request the latency is ", latency)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())
	//Default: r.Use(gin.Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println("example is set by middleware Logger, :", example)
	})

	r.Run(":8080")
}
