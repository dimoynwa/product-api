package generator

import (
	"log"
	"time"
)

var productIdChan = make(chan int)

func InitProductIdGenerator() {
	next := 3
	go func() {
		for ; ; next++ {
			productIdChan <- next
		}
	}()
}

func GenerateProductId() int {
	select {
	case id := <-productIdChan:
		return id
	case <-time.After(2 * time.Second):
		log.Fatalln("Product ID generator is not initialized")
		return 0
	}
}
