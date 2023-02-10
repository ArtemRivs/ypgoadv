package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	fmt.Println("word:", wordPtr)

	// flag.Parse()
	// fmt.Println("word:", wordPtr)
}
