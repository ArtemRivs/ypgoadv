package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("bytes b:", b)
	fmt.Println("string b:", hex.EncodeToString(b))
	// bytes b: [143 10 59 29 240 142 126 163 204 202 61 209 185 219 108 126]
	// string b: 8f0a3b1df08e7ea3ccca3dd1b9db6c7e

}
