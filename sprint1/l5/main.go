package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    currentTimeStr := "2021-09-19T15:59:41+03:00"
    // скопируйте блок себе в IDE и допишите код
    layout := "..."
    currentTime, err := time.Parse(layout, currentTimeStr)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(currentTime)

    now := time.Now()

    fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
    fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
    fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))
}
