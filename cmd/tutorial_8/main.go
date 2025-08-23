package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var results = []int{}
var mu = sync.Mutex{}

func main() {
	ids := [5]int{1, 2, 3, 4, 5}
	start := time.Now()
	for id := range ids {
		wg.Add(1)
		go db_call(ids[id])
	}
	wg.Wait()
	fmt.Println("results", results)
	fmt.Println("Elapsed", time.Since(start))
}

func db_call(id int) {
	time.Sleep(2 * time.Second)
	fmt.Println("DB call for ids", id)
	mu.Lock()
	results = append(results, id)
	mu.Unlock()
	wg.Done()
}
