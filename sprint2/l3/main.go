package main

import (
	"fmt"
	"os"
)

func main() {
	// wordPtr := flag.String("word", "foo", "a string")
	// fmt.Println("word:", wordPtr),
	// flag.Parse()
	// fmt.Println("word:", wordPtr)
	args := os.Args
	fmt.Printf("All args: %v\n", args)
	command := os.Args[0]
	fmt.Printf("Command name: %v\n", command)
	params := os.Args[1:]
	fmt.Printf("Params: %v\n", params)

}
