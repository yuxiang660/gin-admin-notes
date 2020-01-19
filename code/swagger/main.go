package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// User defines the structure of user information from client
type User struct {
	Username string `json:"username"` // Username from client
	Password string `json:"password"` // Password from client
}

func main() {
	router := gin.Default()

	router.OPTIONS("/api", GetOptions)
	router.POST("/api", Register)
	router.GET("/api", GetMessage)

	router.Static("/swagger", "./docs")

	router.Run(":8080")
}

// GetOptions get options for /api url 
// @Tags Methods
// @Summary Get Options
// @Success 200 {string} string	"ok"
// @Router /api [options]
func GetOptions(c *gin.Context) {
	c.Header("Allow", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
}

// Register a user, such as {"username": "root", "password": "123"} 
// @Tags User Management
// @Summary Login
// @Param user body User true "Register a user"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "We need username and password!!"
// @Router /api [post]
func Register(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Username == "" || json.Password == "" {
		c.String(http.StatusBadRequest, "Input username and password")
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are registered"})
}

// GetMessage gets a message from client
// @Tags Message Management
// @Summary Gets a message from client
// @Param message query string false "Message"
// @Success 200 {string} string	"ok"
// @Router /api [get]
func GetMessage(c *gin.Context) {
	message, _ := c.GetQuery("message")
	c.String(http.StatusOK, "Get works! you sent: " + message)
}
