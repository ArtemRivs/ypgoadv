package main

import (
	"fmt"
	"sync"
	"time"
)

type bill struct {
	mu     *sync.Mutex
	amount map[int]int
}

func (b *bill) subAmount(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.amount[1] > 0 {
		time.Sleep(10 * time.Millisecond)
		b.amount[1] -= amount
	}
}

func main() {

	userBill := &bill{
		mu:     &sync.Mutex{},
		amount: map[int]int{1: 1000},
	}

	for i := 0; i < 1000; i++ {
		go userBill.subAmount(10)
	}

	time.Sleep(time.Second)
	fmt.Println(userBill.amount)

}
