package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
)

func getPath(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Get id:", id)

	c.String(200, "Success")
}

func main() {
	// Test with:
	// curl -X GET "localhost:8085/testing/1"
	route := gin.Default()
	route.GET("/testing/:id", getPath)
	route.Run(":8085")
}