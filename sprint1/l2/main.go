package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("Mon 2 Jan 2006 15:04:05 MST"))
}
