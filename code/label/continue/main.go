package main

import "fmt"

func main() {
	fmt.Println("continue label")

	cmds := []string{"build", "invalid", "deploy"}

	OutLoop:
		for _, cmd := range cmds {
			switch cmd {
			case "build":
				fmt.Println("building")
			case "invalid":
				fmt.Println("invalid")
				continue OutLoop
			case "deploy":
				fmt.Println("deploying")
			default:
				fmt.Println("error")
			}
			fmt.Println("success")
		}
}
