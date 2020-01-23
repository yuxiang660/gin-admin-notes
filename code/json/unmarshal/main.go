package main

import (
	"encoding/json"
	"log"
	"fmt"
	"time"
)

type FruitBasket struct {
    Name    string
    Fruit   []string
    ID      int64  `json:"ref"` // Change the tag name from `ID` to `ref`
    private string // Only export public fields
    Created time.Time
}

func main() {
	jsonData := []byte(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999,
		"Created": "2018-04-09T23:00:00Z"
	}`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.ID, basket.Created)

}
