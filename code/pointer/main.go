package main

import "fmt"

type Person struct {
	age int
}

type PersonInfo interface {
	GetAge() int
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) Grow() {
	p.age++
}

func (p Person) GetAge() int {
	return p.age
}

func (p Person) Grow() {
	p.age++
}

func main() {
	fmt.Println("Pointer and Interface")

	var person Person
	var personPointer *Person
	var personInterface PersonInfo
	var personPointerInterface PersonInfo

	person = Person{age: 1}
	personPointer = &Person{age: 1}
	personInterface = person
	personPointerInterface = personPointer

	fmt.Println("person age:", person.GetAge())
	fmt.Println("personPointer age:", personPointer.GetAge())
	fmt.Println("personInterface age:", personInterface.GetAge())
	fmt.Println("personPointerInterface age:", personPointerInterface.GetAge())

	person.Grow()
	personPointer.Grow()

	fmt.Println("After grow:")
	fmt.Println("person age:", person.GetAge())
	fmt.Println("personPointer age:", personPointer.GetAge())
	fmt.Println("personInterface age:", personInterface.GetAge())
	fmt.Println("personPointerInterface age:", personPointerInterface.GetAge())
}

