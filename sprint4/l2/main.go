package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Data struct {
    // определите здесь нужные поля
}

func main() {
    foo := []Data{
        {
            ID:   10,
            Name: "John Doe",
        },
        {
            Name:    "Вася",
            Company: "Яндекс",
        },
    }
    out, err := json.Marshal(foo)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
} 
