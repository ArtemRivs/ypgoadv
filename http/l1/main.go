package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error")
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		fmt.Println(k, ":", v)
	}
	fmt.Println("")
	// fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Cookies:", resp.Cookies())

	// body, err := io.ReadAll(resp.Body)
	// body, err := resp.Header.
	// fmt.Printf("%s", body)

	// for true {
	// 	bs := make([]byte, 1024)
	// 	n, err := resp.Body.Read(bs)
	// 	if err != nil {
	// 		fmt.Println("BS Error")
	// 	}
	// 	fmt.Println(string(bs[:n]))
	// 	if n == 0 || err != nil {
	// 		break
	// 	}
	// }

}
