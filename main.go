package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Execting =>", time.Now())
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	wg.Wait()
}
