package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var ops int64

	fmt.Println(state, mutex)

	for i := 0; i < 100; i++ {
		go func() {
			var total int
			for {
				key := rand.Intn(7)

				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(7)
				val := rand.Intn(100)

				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
