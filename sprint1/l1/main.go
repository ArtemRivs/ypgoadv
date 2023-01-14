package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// timeStr := now.Format("1.2.06 3:4:5 -07 MST")
	timeStr := now.Format("02.01.2006 15:4:5")

	fmt.Println(timeStr)
}
