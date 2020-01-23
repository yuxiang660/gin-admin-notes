package main

import (
	"fmt"
	"encoding/json"
	"log"
	
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
}

func main() {
	fmt.Println("Outpt Buffer")

	users := []User{
		User{Username: "abc", Password: "123"},
		User{Username: "dfe", Password: "456"},
	}
	fmt.Println(users)

	body, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(body)

	router := gin.Default()
	router.GET("", func(c *gin.Context){
		c.Data(200, "application/json; charset=utf-8", body)
		//c.JSON(200, body)
		c.Abort()
	})

	router.Run(":12345")
}
