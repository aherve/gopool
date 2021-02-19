# Go Pool

Simple implementation of a concurrency pool with limited concurrency

## Example usage:

```go
package main

import (
	"log"
	"time"

	"github.com/aherve/gopool"
)

func main() {

	pool := gopool.NewPool(3)
	for i := 0; i < 10; i++ {
		pool.Add(1)
		go work(i, pool)
	}

	pool.Wait()
	log.Println("All Done !")
}

func work(i int, pool *gopool.GoPool) {
	defer pool.Done()
	log.Printf("working hard on %v", i)
	time.Sleep(time.Second)
	log.Printf("%v is done", i)
}
```
