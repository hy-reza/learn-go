package Day5

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{New: func() interface{} {
		return "New"
	}}

	pool.Put("Handy")z
	pool.Put("Reza")
	pool.Put("Alfanda")

	for i := 1; i <= 10; i++ {
		go func() {
			text := pool.Get()
			fmt.Println(text)
			time.Sleep(1 * time.Second)

			pool.Put(text)
		}()
	}

	time.Sleep(5 * time.Second)
}
