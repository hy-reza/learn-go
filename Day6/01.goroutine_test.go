package Day5

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func sayHi() {
	fmt.Println("Hello")
}
func sayBye() {
	fmt.Println("Bye")
}

func TestCreateGoroutine(t *testing.T) {
	go sayHi()
	fmt.Println("Ups")
	go sayBye()

	time.Sleep(1 * time.Second)
}

func displayNum(num int) {
	fmt.Println("num : ", num)
}

func TestDisplayNum(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go displayNum(i)
	}

	time.Sleep(5 * time.Second)
}

func TestRaceChan(t *testing.T) {
	c := 0
	var mutex sync.RWMutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				c += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(c)
}

type Account struct {
	sync.RWMutex
	balance int
}

func (a *Account) changeBalance(m int) {
	a.RWMutex.Lock()
	a.balance += m
	a.RWMutex.Unlock()

}

func (a *Account) getBalance() int {
	a.RWMutex.RLock()
	balance := a.balance
	a.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	myAccount := Account{}

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				myAccount.changeBalance(1)
				fmt.Println(myAccount.getBalance())
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("total balance : ", myAccount.getBalance())
}
