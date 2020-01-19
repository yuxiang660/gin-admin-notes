package main

import (
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
)

func main() {
	fmt.Println("Go Buntdb")

	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("name", "risa", nil)
		return err
	})

	var name string
	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("name")
		if err != nil {
			log.Fatal(err)
			return err
		}
		name = val
		return nil
	})
	fmt.Println(name)
}
