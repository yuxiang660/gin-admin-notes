package main

import "fmt"

func main() {
	fmt.Println("break label")

	cmds := []string{"build", "test", "stop", "deploy"}

OutLoop:
	for _, cmd := range cmds {
		switch cmd {
		case "build":
			fmt.Println("building")
		case "test":
			fmt.Println("testing")
			break // no effect since go switch always break here
		case "stop":
			fmt.Println("stoping")
			break OutLoop
		case "deploy":
			fmt.Println("deploying")
		default:
			fmt.Println("error")
		}
	}
}
