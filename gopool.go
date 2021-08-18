//Package gopool provides Add, Done and Wait methods just like sync.WaitGroup would do.
//On top of that, a max concurrency is applied so that the number of goroutines stays under control
package gopool

// GoPool describe a concurrency pool
type GoPool struct {
	buffer     chan bool
	bufferSize int
}

// NewPool creates a new pool with max concurrency size
func NewPool(maxConcurrency int) *GoPool {
	buffer := make(chan bool, maxConcurrency)
	return &GoPool{
		buffer:     buffer,
		bufferSize: maxConcurrency,
	}
}

// Add declares new tasks
func (gp *GoPool) Add(n int) {
	if n < 0 {
		panic("n cannot be < 0")
	}
	for i := 0; i < n; i++ {
		gp.buffer <- true // take a slot in buffer channel
	}
}

// Done frees a slot
func (gp *GoPool) Done() {
	<-gp.buffer // read a value from the buffer
}

// Wait can be used to wait for all goroutines to finish
func (gp *GoPool) Wait() {
	// insert `maxConcurrency` tasks that do nothing, then clear them all
	for i := 0; i < gp.bufferSize; i++ {
		gp.Add(1)
		defer gp.Done()
	}
}
