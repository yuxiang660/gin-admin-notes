package main

import (
	"fmt"
	"log"
	"strconv"
	"encoding/json"
)

type User struct {
	Username string
	Password string
}

func Serialize(entries []User) map[string]interface{} {
	data := make(map[string]interface{})
	for i, entry := range entries {
		data[strconv.Itoa(i)] = entry
	}
	return data
}

func main() {
	users := []User{
		User{Username: "abc", Password: "123"},
		User{Username: "dfe", Password: "456"},
	}
	fmt.Println(users)

	jsonData, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}
