package main

import (
	"errors"
	"log"
	"fmt"

	"go.uber.org/dig"
)

type Hello struct {
	name string
}

func NewHello() (*Hello, error) {
	err := errors.New("test errror")
	return nil, err
}

func (h *Hello) greet() {
	fmt.Println("hello")
}


func BuildContainer() (*dig.Container, error) {
	c := dig.New()

	err := c.Provide(NewHello)

	return c, err
}

func main() {
	c, err := BuildContainer()
	if err != nil {
		log.Fatal("err1", err)
	}

	err = c.Invoke(func(h *Hello) {
		h.greet()
	})
	if err != nil {
		fmt.Println("errrrrrrrrrrrr")
		log.Fatal(err)
	}
}