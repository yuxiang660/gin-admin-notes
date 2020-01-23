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
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		ID:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	dataPretty, err := json.MarshalIndent(basket, "", "	")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(dataPretty))

}
