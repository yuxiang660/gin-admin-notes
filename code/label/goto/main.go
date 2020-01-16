package main

import "fmt"

func main() {
	fmt.Println("goto label")

	x := 1
	goto Done
	x = 2
Done:
	fmt.Println(x)
}
