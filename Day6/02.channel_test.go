package Day5

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	c := make(chan string)
	defer close(c)

	go func(c chan<- string) { //in<- chan
		time.Sleep(1 * time.Second)
		c <- "handy reza alfanda3"
		fmt.Println("Selesai kirim data")
	}(c)

	go func(name <-chan string) { //<-out chan
		fmt.Println("Hello my name is ", name)
		time.Sleep(5 * time.Second)
	}(c)

}

func TestRangeChan(t *testing.T) {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Value ke: " + strconv.Itoa(i)
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Done gan !")
}

func TestSelectChan(t *testing.T) {
	chan1 := make(chan string)
	chan2 := make(chan string)
	defer close(chan1)
	defer close(chan2)

	go func(c1 chan<- string) {
		c1 <- "channel 1"
	}(chan1)
	go func(c2 chan<- string) {
		c2 <- "channel 2"
	}(chan2)

	count := 0
	for {

		select {
		case data := <-chan1:
			fmt.Println(data)
			count++
		case data := <-chan2:
			fmt.Println(data)
			count++
		default:
			fmt.Println("Please wait ...")
		}

		if count >= 2 {
			break
		}
	}

}

func sayHii(g *sync.WaitGroup, i int) {
	defer g.Done()

	g.Add(1)

	fmt.Println("Hi", i)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	g := &sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		go sayHii(g, i)
	}

	g.Wait()
	fmt.Println("Beres")
}
