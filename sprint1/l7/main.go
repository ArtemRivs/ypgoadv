package main

import (
	"fmt"
	"time"
)

func main() {

	birthday := time.Date(1993+100, time.November, 26, 0, 0, 0, 0, time.Local)
	duration := birthday.Sub(time.Now())
	days := int(duration.Hours() / 24)
	fmt.Println(days)

}
