package main

import (
	// "crypto/rand"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func worker(name string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("Hello from %s line %d\n", name, i)
// 	}
// }

func advancedWorker(name string, c chan string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("Hello from %s line %d\n", name, i)
	}
}

func switchboard(cap int, co, ct, st chan string) {
	counter := 0
	for {
		select {
		case o := <-co:
			fmt.Println(o)
			counter++
			if counter == cap {
				st <- "DONE"
			}
			// break
		case t := <-ct:
			fmt.Println(t)
			counter++
			if counter == cap {
				st <- "DONE"
			}
			// break
		}
	}
}

type Shared struct {
	Data map[int]int
	sync.Mutex
}

func NewShared() *Shared {
	return &Shared{
		Data: make(map[int]int),
	}
}

func main() {
	// block := make(chan bool)

	// // block<- true
	// // go worker("One")
	// // go worker("Two")

	// go func(name string, b chan bool) {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Printf("Hello from %s line %d\n", name, i)
	// 	}
	// 	block <- true
	// }("Three", block)

	// <-block

	// *****************************************************

	// cone := make(chan string)
	// ctwo := make(chan string)
	// stopper := make(chan string)

	// go switchboard(30, cone, ctwo, stopper)

	// go advancedWorker("One", cone)
	// go advancedWorker("Two", ctwo)
	// go advancedWorker("Three", ctwo)

	// <-stopper

	// *****************************************************

	// shared := make(map[int]int)
	// mutex := &sync.Mutex{}

	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		k := rand.Intn(5)
	// 		v := rand.Intn(100)
	// 		mutex.Lock()
	// 		shared[k] = v // race condition
	// 		mutex.Unlock()
	// 	}()
	// }

	// time.Sleep(time.Second)
	// pp.Println(shared)

	// *****************************************************

	// shared := NewShared()

	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		k := rand.Intn(5)
	// 		v := rand.Intn(100)
	// 		shared.Lock()
	// 		shared.Data[k] = v // race condition
	// 		shared.Unlock()
	// 	}()
	// }

	// time.Sleep(time.Second)
	// pp.Println(shared)

	// *****************************************************

	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func(ix int) {
			defer group.Done()
			fmt.Printf("Job %d started\n", ix)
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
			fmt.Printf("Job %d ended\n", ix)
			// ...
		}(i)
	}

	group.Wait()
}

/*

ASSIGNMENT

Create a in-memmory kay-value storage mechanism that deletes records on a previously set time

StoreData("name", "Peter", 10) // set data with key "name" and value "Peter", expire after 10 seconds
StoreData("age", 34, 5) // set data with key "age" and value 34, expire in 5 sedonds

// 7 seconds passed
val, err := GetData("name") // val = Peter // retrieve data with key "name" from store, if not expired
val2, err := GetData("age") // val = nil, err = ERROR! (data removed or non-existing) // retrieve data with key "age" from store, if not expired

*/
