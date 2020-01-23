package main

import (
	"fmt"
)

type Inner struct {
	ID int
}

func (i *Inner) getID() int {
	return i.ID
}

type Outer struct {
	Inner
	Name string
}

func main() {
	fmt.Println("Struct Test")

	var o Outer
	o.ID = 1
	o.Inner.ID = 2
	o.Name = "abc"

	fmt.Println(o)
	fmt.Println("getID():", o.getID())
}