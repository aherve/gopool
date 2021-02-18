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

// Add immediately declares a task, but wait for a slot to return
func (gp *GoPool) Add(i int) {
	gp.wg.Add(i)
	gp.sem <- true // take a slot in semaphore channel
}

// Done frees a slot
func (gp *GoPool) Done() {
	<-gp.sem // free a slot
	gp.wg.Done()
}

// Wait can be used to wait for all goroutines to finish
func (gp *GoPool) Wait() {
	gp.wg.Wait()
}
