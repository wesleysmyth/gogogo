package main

import (
	"fmt"
	"sync"

	"golang.org/x/exp/rand"
)

func generator(intChannel chan int) {
	for {
		intChannel <- rand.Intn(100) + 1
	}
}

func reader(intChannel chan int) {
	for {
		fmt.Println(<-intChannel)
	}
}

func main() {
	fmt.Println("Hello, World!")
	intChannel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		generator(intChannel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		reader(intChannel)
	}()

	wg.Wait()
}
