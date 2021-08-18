package gopool_test

import (
	"log"
	"time"

	"github.com/aherve/gopool"
)

func Example() {

	pool := gopool.NewPool(3) // creates pool with limited concurrency of 3
	for i := 0; i < 10; i++ {
		pool.Add(1) // This will pause until a slot is available
		go work(i, pool)
	}

	pool.Wait()
	log.Println("All Done !")
}

func work(i int, pool *gopool.GoPool) {
	defer pool.Done() // just like with sync.WaitGroup
	log.Printf("working hard on %v", i)
	time.Sleep(time.Second)
	log.Printf("%v is done", i)
}
