package main

import (
	"fmt"
	"log"
	"encoding/json"
)

func main() {
	jsonData := []byte(`
	{
		"Name": "Eve",
		"Age": 6,
		"Parents": [
			"Alice",
			"Bob"
		]
	}`)
	
	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	for k, v := range data {
		fmt.Println(k, v)
	}

	data2 := map[string]interface{}{
		"Name": "Eve",
		"Age": 6,
		"Parents": []string{"Alice", "Bob"},
	}
	jsonData2, err := json.Marshal(data2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData2))

}