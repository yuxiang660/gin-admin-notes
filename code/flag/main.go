package main

import (
	"fmt"
	"flag"
)

// 1. go run main.go -h
// 2. go run main.go -word=opt -numb=7 -fork -svar=flag
// 3. go run main.go -word=opt
// 4. go run main.go -word=opt a1 a2 a3
// 5. go run main.go -word=opt a1 a2 a3 -numb=7
// 6. go run main.go -wat
func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}