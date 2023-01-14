package main

import (
	"fmt"
	"time"
)

func main() {

//     time.AfterFunc(1*time.Second, func() {
//         fmt.Println("hi from AfterFunc")
//     })
//     fmt.Println("hi")
//     time.Sleep(2 * time.Second)
//     fmt.Println("goodbye")
	
//     start := time.Now()
//     timer := time.NewTimer(2 * time.Second) // создаём таймер
//     t := <-timer.C                          // ожидаем срабатывания таймера
//     fmt.Println(t.Sub(start).Seconds())     // выводим разницу во времени	

    start := time.Now()
    ticker := time.NewTicker(2 * time.Second)
    for i := 0; i < 10; i++ {
        t := <-ticker.C
        fmt.Println(int(t.Sub(start).Seconds()))
    }
	
	
}
