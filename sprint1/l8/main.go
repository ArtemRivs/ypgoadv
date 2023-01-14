package main

import (
	"fmt"
	"time"
)

func main() {

    time.AfterFunc(1*time.Second, func() {
        fmt.Println("hi from AfterFunc")
    })
    fmt.Println("hi")
    time.Sleep(2 * time.Second)
    fmt.Println("goodbye")

}
