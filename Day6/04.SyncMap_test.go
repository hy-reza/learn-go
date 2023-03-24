package Day5

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	// Membuat sebuah sync map
	var m sync.Map

	// Menambahkan data ke sync map secara asynchronous dengan go routine
	go func() {
		defer wg.Done()
		m.Store("name", "John Doe")
		m.Store("age", 30)
		m.Store("address", "123 Main St")
	}()

	// Mengambil data dari sync map secara asynchronous dengan go routine
	go func() {
		defer wg.Done()
		if v, ok := m.Load("name"); ok {
			fmt.Println("Name:", v)
		}
		if v, ok := m.Load("age"); ok {
			fmt.Println("Age:", v)
		}
		if v, ok := m.Load("address"); ok {
			fmt.Println("Address:", v)
		}
	}()

	wg.Wait()
}
