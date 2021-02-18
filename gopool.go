package gopool

import "sync"

//GoPool provides Add, Done and Wait methods just like sync.WaitGroup
//On top of that, a max concurrency is applied so that the number of goroutines stays under control
type GoPool struct {
	wg  *sync.WaitGroup
	sem chan bool
}

// NewPool creates a new pool with max concurrency size
func NewPool(maxConcurrency int) *GoPool {
	var wg sync.WaitGroup
	sem := make(chan bool, maxConcurrency)
	return &GoPool{
		wg:  &wg,
		sem: sem,
	}
}

// Add waits for a slot to be available, then calls waitGroup.Add
func (gp *GoPool) Add(i int) {
	gp.sem <- true // take a slot in semaphore channel
	gp.wg.Add(i)
}

// Done frees a slot
func (gp *GoPool) Done() {
	gp.wg.Done()
	<-gp.sem // free a slot
}

// Wait can be used to wait for all goroutines to finish
func (gp *GoPool) Wait() {
	gp.wg.Wait()
}
