package Day5

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var count int64 = 0
	// mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&count, 1)
			}

		}()
	}

	wg.Wait()
	fmt.Println("final result : ", count)
}

func TestGOMAXPROCS(t *testing.T) {

	totalCPU := runtime.NumCPU()
	fmt.Println("CPU : ", totalCPU)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread : ", totalThread)
	totalRoutine := runtime.NumGoroutine()
	fmt.Println("Routine : ", totalRoutine)
}
